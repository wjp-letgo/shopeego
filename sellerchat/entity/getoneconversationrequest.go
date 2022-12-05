package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetOneConversationRequest struct {
	ConversationId int64 `json:"conversation_id"`
}

func (g GetOneConversationRequest) String() string {
	return lib.ObjectToString(g)
}
