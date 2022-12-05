package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetAddOnDealMainItemResponseResponseEntity struct {
	MainItemList []GetAddOnDealMainItemMainItemListResponseEntity `json:"main_item_list"`
	AddOnDealId  int64                                            `json:"add_on_deal_id"`
}

func (g GetAddOnDealMainItemResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
