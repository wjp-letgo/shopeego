package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetConversationListResult struct {
	Error     string                                    `json:"error"`
	Message   string                                    `json:"message"`
	RequestId string                                    `json:"request_id"`
	Response  GetConversationListResponseResponseEntity `json:"response"`
}

func (g GetConversationListResult) String() string {
	return lib.ObjectToString(g)
}
