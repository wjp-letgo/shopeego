package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteConversationRequest struct {
	ConversationId int64 `json:"conversation_id"`
}

func (g DeleteConversationRequest) String() string {
	return lib.ObjectToString(g)
}
