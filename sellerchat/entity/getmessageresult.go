package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetMessageResult struct {
	Error     string                           `json:"error"`
	Message   string                           `json:"message"`
	RequestId string                           `json:"request_id"`
	Response  GetMessageResponseResponseEntity `json:"response"`
}

func (g GetMessageResult) String() string {
	return lib.ObjectToString(g)
}
