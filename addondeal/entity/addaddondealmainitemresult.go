package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type AddAddOnDealMainItemResult struct {
	Error     string                                     `json:"error"`
	Message   string                                     `json:"message"`
	RequestId string                                     `json:"request_id"`
	Response  AddAddOnDealMainItemResponseResponseEntity `json:"response"`
}

func (g AddAddOnDealMainItemResult) String() string {
	return lib.ObjectToString(g)
}
