package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateDiscountItemErrorListResponseEntity struct {
	ItemId      int64  `json:"item_id"`
	ModelId     int64  `json:"model_id"`
	FailMessage string `json:"fail_message"`
	FailError   string `json:"fail_error"`
}

func (g UpdateDiscountItemErrorListResponseEntity) String() string {
	return lib.ObjectToString(g)
}
