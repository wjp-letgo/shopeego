package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetConversationListPageResultResponseEntity struct {
	PageSize   int                                         `json:"page_size"`
	NextCursor GetConversationListNextCursorResponseEntity `json:"next_cursor"`
	More       bool                                        `json:"more"`
}

func (g GetConversationListPageResultResponseEntity) String() string {
	return lib.ObjectToString(g)
}
