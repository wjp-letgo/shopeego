package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateAddOnDealResponseResponseEntity struct {
	StartTime              int     `json:"start_time"`
	EndTime                int     `json:"end_time"`
	PromotionType          int     `json:"promotion_type"`
	PurchaseMinSpend       float32 `json:"purchase_min_spend"`
	AddOnDealId            int64   `json:"add_on_deal_id"`
	PerGiftNum             int     `json:"per_gift_num"`
	PromotionPurchaseLimit []int   `json:"promotion_purchase_limit"`
	AddOnDealName          string  `json:"add_on_deal_name"`
}

func (g UpdateAddOnDealResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
