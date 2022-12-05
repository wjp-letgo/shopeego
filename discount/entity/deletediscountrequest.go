package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteDiscountRequest struct {
	DiscountId int64 `json:"discount_id"`
}

func (g DeleteDiscountRequest) String() string {
	return lib.ObjectToString(g)
}
