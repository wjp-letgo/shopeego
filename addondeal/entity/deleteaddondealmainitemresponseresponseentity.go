package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteAddOnDealMainItemResponseResponseEntity struct {
	MainItemList []int `json:"main_item_list"`
	AddOnDealId  int64 `json:"add_on_deal_id"`
}

func (g DeleteAddOnDealMainItemResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
