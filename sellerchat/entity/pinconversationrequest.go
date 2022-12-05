package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type PinConversationRequest struct {
	ConversationId int64 `json:"conversation_id"`
}

func (g PinConversationRequest) String() string {
	return lib.ObjectToString(g)
}
