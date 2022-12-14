package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type EndDiscountResponseResponseEntity struct {
	DiscountId int64 `json:"discount_id"`
	ModifyTime int   `json:"modify_time"`
}

func (g EndDiscountResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
