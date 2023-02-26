package taikang

import (
	"context"
	"encoding/json"
	"fmt"
	"gitlab.mycool.tv/local/helper"
	"log"
	"time"
)

var (
	punchChan = make(chan *Punch, 20)
)

// Punch 运动打卡信息
type Punch struct {
	HotelId     int    `json:"hotelId"`
	GuestNumber string `json:"guestNumber"`
	GuestName   string `json:"GuestName"`
	RoomNumber  string `json:"roomNumber"`
	StartTime   int64  `json:"startTime"`
	EndTime     int64  `json:"endTime"`
}

// PunchRoutine 运动健康打卡goroutine
type PunchRoutine struct {
	url      string
	interval time.Duration
	punchQ   []*Punch
}

func NewPunchRoutine(url string, interval time.Duration) *PunchRoutine {
	return &PunchRoutine{
		url:      url,
		interval: interval,
		punchQ:   make([]*Punch, 0),
	}
}

func (routine *PunchRoutine) Start(ctx context.Context) {
	ticker := time.NewTicker(routine.interval)
	defer ticker.Stop()
	for {
		select {
		case punch := <-punchChan:
			routine.punchQ = append(routine.punchQ, punch)
		case <-ticker.C:
			routine.post()
		case <-ctx.Done():
			routine.post()
			log.Printf("exit taikang punch routine: %s", ctx.Err())
			return
		}
	}
}

// post 请求，这里如果一直失败，punchQ会一直增加，最终OOM
func (routine *PunchRoutine) post() {
	if len(routine.punchQ) == 0 {
		log.Print("no taikang punch to post")
		return
	}
	hotelId := routine.punchQ[0].HotelId
	posturl := fmt.Sprintf("%s?hotelId=%d", routine.url, hotelId)

	p, _ := json.Marshal(routine.punchQ)
	log.Printf("post punch: %s", p)

	p, e := helper.HTTPPostJSON(posturl, string(p))
	m := new(helper.Msg)

	if e = helper.JSONResponse(p, e, m); e != nil {
		log.Printf("post punch err: %s", e)
		return
	}

	if m.Code != helper.CodeSuccess {
		log.Printf("post punch err: %v", m)
		return
	}

	log.Printf("post punch success: post size = %d", len(routine.punchQ))
	routine.punchQ = routine.punchQ[:0]
}
