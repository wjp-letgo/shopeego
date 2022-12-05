package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateDiscountItemResponseResponseEntity struct {
	DiscountId int64                                       `json:"discount_id"`
	Count      int                                         `json:"count"`
	ErrorList  []UpdateDiscountItemErrorListResponseEntity `json:"error_list"`
}

func (g UpdateDiscountItemResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
