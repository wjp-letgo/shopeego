package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type ReadConversationRequest struct {
	ConversationId    int64  `json:"conversation_id"`
	LastReadMessageId string `json:"last_read_message_id"`
}

func (g ReadConversationRequest) String() string {
	return lib.ObjectToString(g)
}
