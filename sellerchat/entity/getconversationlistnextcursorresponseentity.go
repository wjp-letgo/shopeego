package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetConversationListNextCursorResponseEntity struct {
	NextMessageTimeNano string `json:"next_message_time_nano"`
	ConversationId      string `json:"conversation_id"`
}

func (g GetConversationListNextCursorResponseEntity) String() string {
	return lib.ObjectToString(g)
}
