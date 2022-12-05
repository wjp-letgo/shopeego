package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type AddAddOnDealSubItemRequest struct {
	AddOnDealId int64                                         `json:"add_on_deal_id"`
	SubItemList []AddAddOnDealSubItemSubItemListRequestEntity `json:"sub_item_list"`
}

func (g AddAddOnDealSubItemRequest) String() string {
	return lib.ObjectToString(g)
}
