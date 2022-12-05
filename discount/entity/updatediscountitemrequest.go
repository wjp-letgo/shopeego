package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateDiscountItemRequest struct {
	DiscountId int64                                     `json:"discount_id"`
	ItemList   []UpdateDiscountItemItemListRequestEntity `json:"item_list"`
}

func (g UpdateDiscountItemRequest) String() string {
	return lib.ObjectToString(g)
}
