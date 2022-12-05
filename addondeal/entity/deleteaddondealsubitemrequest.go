package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteAddOnDealSubItemRequest struct {
	AddOnDealId int64                                            `json:"add_on_deal_id"`
	SubItemList []DeleteAddOnDealSubItemSubItemListRequestEntity `json:"sub_item_list"`
}

func (g DeleteAddOnDealSubItemRequest) String() string {
	return lib.ObjectToString(g)
}
