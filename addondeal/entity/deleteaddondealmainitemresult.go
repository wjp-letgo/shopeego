package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteAddOnDealMainItemResult struct {
	Error     string                                        `json:"error"`
	Message   string                                        `json:"message"`
	RequestId string                                        `json:"request_id"`
	Response  DeleteAddOnDealMainItemResponseResponseEntity `json:"response"`
}

func (g DeleteAddOnDealMainItemResult) String() string {
	return lib.ObjectToString(g)
}
