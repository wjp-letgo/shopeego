package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type SetPushConfigResult struct {
	Status    string `json:"status"`
	RequestId string `json:"request_id"`
	Error     string `json:"error"`
	Message   string `json:"message"`
}

func (g SetPushConfigResult) String() string {
	return lib.ObjectToString(g)
}
