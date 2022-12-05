package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateBundleDealItemResult struct {
	Error     string                                     `json:"error"`
	Message   string                                     `json:"message"`
	RequestId string                                     `json:"request_id"`
	Response  UpdateBundleDealItemResponseResponseEntity `json:"response"`
}

func (g UpdateBundleDealItemResult) String() string {
	return lib.ObjectToString(g)
}
