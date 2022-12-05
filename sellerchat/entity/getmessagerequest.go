package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetMessageRequest struct {
	Offset         string `json:"offset"`
	PageSize       int    `json:"page_size"`
	ConversationId int64  `json:"conversation_id"`
}

func (g GetMessageRequest) String() string {
	return lib.ObjectToString(g)
}
