package policy_type

type PolicyType int32

const (
	Policy_MIN PolicyType = 0
	Policy_MAX PolicyType = 1
	Policy_MID PolicyType = 2
	Policy_AVG PolicyType = 3
)

func (p PolicyType) String() string {
	switch p {
	case Policy_MIN:
		return "MIN"
	case Policy_MAX:
		return "MAX"
	case Policy_MID:
		return "MID"
	case Policy_AVG:
		return "AVG"
	default:
		return "UNKNOWN"
	}
}

type CutStatus int

// 显式类型转换
const (
	NotNotified = CutStatus(0)
	Notified    = CutStatus(1)
)
