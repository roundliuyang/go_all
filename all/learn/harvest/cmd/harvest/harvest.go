package main

import (
	"all/all/learn/harvest/cfg"
	"all/all/learn/harvest/pingback"
	"all/all/learn/harvest/taikang"
	"gitlab.mycool.tv/local/helper"
	"log"
	"net/http"
)

type pbHandler struct {
	analyzers []pingback.Analyzer
}

// 构造器，惨绝人寰
func newPbHandler() *pbHandler {
	h := &pbHandler{
		analyzers: make([]pingback.Analyzer, 0),
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

	pb := pingback.FromRequest(r)
	if pb == nil {
		log.Printf("analyze %s", pb.JSON())
	}

	for _, ana := range h.analyzers {
		tmp := *pb
		ana.Analyze(&tmp)
	}
	m.Code, m.Msg = helper.CodeSuccess, "pingback"
}
