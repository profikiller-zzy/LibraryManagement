package custom_type

import "github.com/goccy/go-json"

type Status int

const (
	Free   = 1 // 空闲
	OnLoan = 2 // 被借阅
)

func (s Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Status) String() string {
	switch s {
	case Free:
		return "空闲"
	case OnLoan:
		return "被借阅"
	default:
		return "空闲"
	}
}
