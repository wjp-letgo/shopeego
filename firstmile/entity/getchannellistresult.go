package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetChannelListResult struct {
	RequestId string                               `json:"request_id"`
	Error     string                               `json:"error"`
	Message   string                               `json:"message"`
	Response  GetChannelListResponseResponseEntity `json:"response"`
}

func (g GetChannelListResult) String() string {
	return lib.ObjectToString(g)
}
