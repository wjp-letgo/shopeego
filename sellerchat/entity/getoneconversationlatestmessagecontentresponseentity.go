package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetOneConversationLatestMessageContentResponseEntity struct {
	Text string `json:"text"`
}

func (g GetOneConversationLatestMessageContentResponseEntity) String() string {
	return lib.ObjectToString(g)
}
