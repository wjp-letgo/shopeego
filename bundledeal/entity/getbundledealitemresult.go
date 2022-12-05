package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetBundleDealItemResult struct {
	Error     string                                  `json:"error"`
	Message   string                                  `json:"message"`
	RequestId string                                  `json:"request_id"`
	Response  GetBundleDealItemResponseResponseEntity `json:"response"`
}

func (g GetBundleDealItemResult) String() string {
	return lib.ObjectToString(g)
}
