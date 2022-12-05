package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetWaybillResult struct {
	RequestId string `json:"request_id"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	File      []byte
}

func (g GetWaybillResult) String() string {
	return lib.ObjectToString(g)
}
