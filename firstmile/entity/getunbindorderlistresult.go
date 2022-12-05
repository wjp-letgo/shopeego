package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetUnbindOrderListResult struct {
	RequestId string                                   `json:"request_id"`
	Error     string                                   `json:"error"`
	Message   string                                   `json:"message"`
	Response  GetUnbindOrderListResponseResponseEntity `json:"response"`
}

func (g GetUnbindOrderListResult) String() string {
	return lib.ObjectToString(g)
}
