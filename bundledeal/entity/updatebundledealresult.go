package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateBundleDealResult struct {
	Error     string                                 `json:"error"`
	Message   string                                 `json:"message"`
	RequestId string                                 `json:"request_id"`
	Response  UpdateBundleDealResponseResponseEntity `json:"response"`
}

func (g UpdateBundleDealResult) String() string {
	return lib.ObjectToString(g)
}
