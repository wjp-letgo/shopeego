package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateDiscountResult struct {
	Error     string                               `json:"error"`
	RequestId string                               `json:"request_id"`
	Response  UpdateDiscountResponseResponseEntity `json:"response"`
	Message   string                               `json:"message"`
}

func (g UpdateDiscountResult) String() string {
	return lib.ObjectToString(g)
}
