package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteAddOnDealMainItemRequest struct {
	AddOnDealId  int64 `json:"add_on_deal_id"`
	MainItemList []int `json:"main_item_list"`
}

func (g DeleteAddOnDealMainItemRequest) String() string {
	return lib.ObjectToString(g)
}
