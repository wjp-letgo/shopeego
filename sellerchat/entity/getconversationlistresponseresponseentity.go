package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetConversationListResponseResponseEntity struct {
	PageResult    GetConversationListPageResultResponseEntity      `json:"page_result"`
	Conversations []GetConversationListConversationsResponseEntity `json:"conversations"`
}

func (g GetConversationListResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
