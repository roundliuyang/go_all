package cancel_by_close

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCancel(t *testing.T) {
	punchOrderExecute := &PunchOrderExecute{
		TaskId:       "joegmvogmpgmp",
		BeginTime:    "gvmnogvmrogvrmovg",
		EndTime:      "nmonvgonfro",
		TotalMinutes: 1,
	}
	p, _ := json.Marshal(punchOrderExecute)
	log.Printf("post yangkangonline request body: %s", string(p))
}

// PunchOrderExecute 运动处方任务信息
type PunchOrderExecute struct {
	TaskId       string
	BeginTime    string
	EndTime      string
	TotalMinutes int64
}
