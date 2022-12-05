package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetDiscountResult struct {
	Error     string                            `json:"error"`
	Message   string                            `json:"message"`
	RequestId string                            `json:"request_id"`
	Response  GetDiscountResponseResponseEntity `json:"response"`
}

func (g GetDiscountResult) String() string {
	return lib.ObjectToString(g)
}
