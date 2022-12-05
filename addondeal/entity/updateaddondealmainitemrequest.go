package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateAddOnDealMainItemRequest struct {
	AddOnDealId  int64                                              `json:"add_on_deal_id"`
	MainItemList []UpdateAddOnDealMainItemMainItemListRequestEntity `json:"main_item_list"`
}

func (g UpdateAddOnDealMainItemRequest) String() string {
	return lib.ObjectToString(g)
}
