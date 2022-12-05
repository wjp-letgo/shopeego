package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type AddDiscountItemResponseResponseEntity struct {
	DiscountId int64                                    `json:"discount_id"`
	Count      int                                      `json:"count"`
	ErrorList  []AddDiscountItemErrorListResponseEntity `json:"error_list"`
}

func (g AddDiscountItemResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
