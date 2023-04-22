package custom_type

import "encoding/json"

type Gender int

const (
	Male   = 1 // 男性
	Female = 2 // 女性
)

func (g Gender) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

func (g Gender) String() string {
	switch g {
	case Male:
		return "男性"
	case Female:
		return "女性"
	default:
		return "男性"
	}
}
