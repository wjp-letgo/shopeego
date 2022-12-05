package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateAddOnDealMainItemResult struct {
	Error       string                                        `json:"error"`
	Message     string                                        `json:"message"`
	RequestId   string                                        `json:"request_id"`
	Response    UpdateAddOnDealMainItemResponseResponseEntity `json:"response"`
	AddOnDealId int64                                         `json:"add_on_deal_id"`
}

func (g UpdateAddOnDealMainItemResult) String() string {
	return lib.ObjectToString(g)
}
