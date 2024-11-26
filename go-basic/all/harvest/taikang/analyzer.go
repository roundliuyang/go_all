package taikang

import (
	pingback2 "all/all/harvest/pingback"
	"time"
)

var (
	maxBeatInterval = 65 * time.Second
)

var _ interface {
	pingback2.Analyzer
	pingback2.Filter
} = (*PunchAnalyzer)(nil)

// 运动健康打卡解析
type PunchAnalyzer struct {
	*pingback2.SimpleAnalyzer
	uid string
}

func NewTaikangPunchAnalyzer(uid string) *PunchAnalyzer {
	ana := &PunchAnalyzer{
		SimpleAnalyzer: &pingback2.SimpleAnalyzer{
			Testers:     make([]pingback2.Filter, 1),
			NewConsumer: func() pingback2.Consumer { return new(pair) },
		},
		uid: uid,
	}
	ana.Testers[0] = ana
	return ana
}

func (analyzer *PunchAnalyzer) Filter(pb *pingback2.Pingback) bool {
	return pb.Uid == analyzer.uid || pb.Action == pingback2.ActionBeat
}

type pair struct {
	in, lastBeat *pingback2.Pingback
}

func (p *pair) Consume(pb *pingback2.Pingback) {
	switch pb.Action {
	case pingback2.ActionPlay:
		// 记录in ,连续多次的in,会刷新成最新的
		p.in = pb
	case pingback2.ActionSkip, pingback2.ActionEnd:
		if p.in != nil {
			// 有 in 有 out ,生成一条行为记录
			punchChan <- &Punch{
				HotelId:     p.in.HotelId,
				GuestNumber: p.in.CurGuestNumber,
				GuestName:   p.in.CurGuestName,
				RoomNumber:  p.in.RoomNumber,
				StartTime:   p.in.Time,
				EndTime:     pb.Time,
			}

			// in 置空，防止连续多次out
			p.in = nil
		}
	case pingback2.ActionBeat:
		if p.lastBeat != nil && pb.Time-p.lastBeat.Time > int64(maxBeatInterval) {
			// 两次心跳超过了最大时间间隔，认为关机了
			if p.in != nil && p.in.Time < p.lastBeat.Time {
				// 未退出直接关机了
				punchChan <- &Punch{
					HotelId:     p.in.HotelId,
					GuestNumber: p.in.CurGuestNumber,
					GuestName:   p.in.CurGuestName,
					RoomNumber:  p.in.RoomNumber,
					StartTime:   p.in.Time,
					EndTime:     p.lastBeat.Time,
				}
				p.lastBeat = nil
			}
			p.lastBeat = nil
		} else {
			p.lastBeat = pb
		}
	}
}
