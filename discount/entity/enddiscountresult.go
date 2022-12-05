package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type EndDiscountResult struct {
	Error     string                            `json:"error"`
	Message   string                            `json:"message"`
	RequestId string                            `json:"request_id"`
	Response  EndDiscountResponseResponseEntity `json:"response"`
}

func (g EndDiscountResult) String() string {
	return lib.ObjectToString(g)
}
