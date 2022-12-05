package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type EndDiscountRequest struct {
	DiscountId int64 `json:"discount_id"`
}

func (g EndDiscountRequest) String() string {
	return lib.ObjectToString(g)
}
