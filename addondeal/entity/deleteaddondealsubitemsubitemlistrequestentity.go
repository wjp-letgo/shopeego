package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteAddOnDealSubItemSubItemListRequestEntity struct {
	ItemId  int64 `json:"item_id"`
	ModelId int64 `json:"model_id"`
}

func (g DeleteAddOnDealSubItemSubItemListRequestEntity) String() string {
	return lib.ObjectToString(g)
}
