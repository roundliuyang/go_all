package pingback

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// action
const (
	ActionIn           = "in"
	ActionOut          = "out"
	ActionPlay         = "play"
	ActionPause        = "pause"
	ActionResume       = "resume"
	ActionSkip         = "skip"
	ActionEnd          = "end"
	ActionBeat         = "beat"
	ActionEsc          = "esc"
	ActionBack         = "back"
	ActionFullscreen   = "fullscreen"
	ActionNofullscreen = "nofullscreen"
	ActionLeft         = "left"
	ActionRight        = "right"
	ActionFastbackward = "fastbackward"
	ActionFastforward  = "fastforward"
	ActionUp           = "up"
	ActionDown         = "down"
	ActionLike         = "like"
	ActionCry          = "cry"
	ActionLove         = "love"
	ActionLaugh        = "laugh"
	ActionShock        = "shock"
	ActionIncr         = "incr"
	ActionDecr         = "decr"
	ActionScanqrcode   = "scanqrcode"
)

type Pingback struct {
	TvIP       string `json:"tvIP"`
	Time       int64  `json:"time"`
	Raw        string `json:"raw"`
	HotelId    int    `json:"hotelId"`
	TemplateId int    `json:"templateId"`
	PageNum    int    `json:"pageNum"`
	LangId     int    `json:"langId"`
	GI         int    `json:"gi"`
	SN         string `json:"sn"`
	RoomNumber string `json:"roomNumber"`
	Item       string `json:"item"`
	ItemName   string `json:"itemName"`
	BId        string `json:"bId"`
	Action     string `json:"action"`
	Type       string `json:"type"`
	GLs        string `json:"gLs"`
	GTs        string `json:"gTs"`
	Gv         string `json:"gv"`
	GBirth     string `json:"gBirth"`
	Genders    string `json:"genders"`
	UserAgent  string `json:"userAgent"`
	Version    string `json:"version"`
	N          string `json:"n"`
	UpItem     string `json:"upItem"`
	CooSource  string `json:"cooSource"`
	UiSource   string `json:"uiSource"`
	GId        string `json:"gId"`
	Cv         string `json:"cv"`

	/* v3 add */
	Uid string `json:"uid"`
	V   string `json:"v"`

	// 泰康运动健康打卡
	CurGuestNumber string `json:"curgNum"`
	CurGuestName   string `json:"curgName"`
}

// FromRequest 从request 中解析pingback
func FromRequest(r *http.Request) *Pingback {
	if r == nil {
		log.Printf("nil request")
		return nil
	}
	xRealIp := r.Header.Get("X-Real-IP")
	if xRealIp == "" {
		xRealIp = r.RemoteAddr
	}
	pb := &Pingback{
		TvIP: xRealIp,
		Time: time.Now().Unix(),
		Raw:  r.RequestURI,
	}
	log.Println(time.Now().Unix())
	combineURL(r.URL, pb)
	return pb
}

// combineURL 解析url query 到pingback
func combineURL(_url *url.URL, pb *Pingback) {
	values := _url.Query()
	pb.Version = strings.Split(_url.Path, "/")[3]

	if !strings.Contains(pb.Version, "v") {
		pb.Version = values.Get("v")
	}
	pb.HotelId, _ = strconv.Atoi(values.Get("hotelId"))
	pb.LangId, _ = strconv.Atoi(values.Get("langId"))
	pb.GI, _ = strconv.Atoi(values.Get("gi"))

	arr := strings.Split(values.Get("item"), "P")
	if len(arr) == 2 {
		pb.PageNum, _ = strconv.Atoi(arr[1])
	}
	pb.Item = arr[0]

	pb.SN = values.Get("sn")
	pb.RoomNumber = values.Get("roomNumber")
	pb.ItemName = values.Get("itemName")
	pb.BId = values.Get("bId")
	pb.Action = values.Get("action")
	pb.Type = values.Get("type")
	pb.GLs = values.Get("gLs")
	pb.GTs = values.Get("gTs")
	pb.Gv = values.Get("gv")
	pb.GBirth = values.Get("gBirth")
	pb.Genders = values.Get("genders")
	pb.N = values.Get("n")
	pb.UpItem = values.Get("upItem")
	pb.CooSource = values.Get("cooSource")
	pb.UiSource = values.Get("uiSource")
	pb.GId = values.Get("gId")
	pb.Cv = values.Get("cv")
	pb.Uid = values.Get("uid")
	pb.V = values.Get("v")

	pb.CurGuestNumber = values.Get("curgNum")
	pb.CurGuestName = values.Get("curgName")
}

func (pb *Pingback) JSON() string {
	p, _ := json.Marshal(pb)
	return string(p)
}
