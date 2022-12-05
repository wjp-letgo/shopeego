package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateDiscountItemModelListRequestEntity struct {
	ModelId             int64   `json:"model_id"`
	ModelPromotionPrice float32 `json:"model_promotion_price"`
}

func (g UpdateDiscountItemModelListRequestEntity) String() string {
	return lib.ObjectToString(g)
}
