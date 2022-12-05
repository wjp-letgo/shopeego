package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteBundleDealResult struct {
	Error     string                                 `json:"error"`
	Message   string                                 `json:"message"`
	RequestId string                                 `json:"request_id"`
	Response  DeleteBundleDealResponseResponseEntity `json:"response"`
}

func (g DeleteBundleDealResult) String() string {
	return lib.ObjectToString(g)
}
