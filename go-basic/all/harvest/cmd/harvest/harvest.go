package main

import (
	"all/all/harvest/cfg"
	pingback2 "all/all/harvest/pingback"
	"all/all/harvest/taikang"
	"gitlab.mycool.tv/local/helper"
	"log"
	"net/http"
)

func pingbackRecordHandler(w http.ResponseWriter, r *http.Request) {
	//m := helper.NewMsg()
	//defer helper.WriteMsg(w, m)
	//
	//var (
	//	e          error
	//	item       = r.FormValue("item")
	//	roomNumber = r.FormValue("roomNumber")
	//	sTime      = r.FormValue("sTime")
	//	eTime      = r.FormValue("eTime")
	//)
	//if item == "" {
	//	m.Code, m.Msg = helper.CodeError, "require item"
	//	return
	//}
	//if roomNumber == "" {
	//	m.Code, m.Msg = helper.CodeError, "require roomNumber"
	//	return
	//}
	//st, e := strconv.ParseInt(sTime, 10, 64)
	//if e != nil {
	//	m.Code, m.Msg = helper.CodeError, "require sTime"
	//	return
	//}
	//et, e := strconv.ParseInt(eTime, 10, 64)
	//if e != nil {
	//	m.Code, m.Msg = helper.CodeError, "require eTime"
	//	return
	//}
	//
	//sT := time.Unix(st, 0)
	//eT := time.Unix(et, 0)

}

type pbHandler struct {
	analyzers []pingback2.Analyzer
}

// 构造器，惨绝人寰
func newPbHandler() *pbHandler {
	h := &pbHandler{
		analyzers: make([]pingback2.Analyzer, 0),
	}
	if cfg.Cfg.Media.Enabled {
		//todo
		//append(h.analyzers,pingback.)
	}
	if cfg.Cfg.Taikang.Punch.Enabled {
		h.analyzers = append(h.analyzers, taikang.NewTaikangPunchAnalyzer(cfg.Cfg.Taikang.Punch.Uid))
	}
	return h
}

func (h pbHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := helper.NewMsg()

	defer helper.WriteMsg(w, m)

	pb := pingback2.FromRequest(r)
	if pb == nil {
		log.Printf("analyze %s", pb.JSON())
	}

	for _, ana := range h.analyzers {
		tmp := *pb
		ana.Analyze(&tmp)
	}
	m.Code, m.Msg = helper.CodeSuccess, "pingback"
}
