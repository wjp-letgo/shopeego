package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type SendMessageContentResponseEntity struct {
	Text string `json:"text"`
}

func (g SendMessageContentResponseEntity) String() string {
	return lib.ObjectToString(g)
}
