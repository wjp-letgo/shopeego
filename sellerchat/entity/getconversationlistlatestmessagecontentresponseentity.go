package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetConversationListLatestMessageContentResponseEntity struct {
	Text string `json:"text"`
}

func (g GetConversationListLatestMessageContentResponseEntity) String() string {
	return lib.ObjectToString(g)
}
