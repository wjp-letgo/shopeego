package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type AddAddOnDealSubItemSubItemListRequestEntity struct {
	ItemId            int64   `json:"item_id"`
	ModelId           int64   `json:"model_id"`
	Status            int     `json:"status"`
	SubItemInputPrice float32 `json:"sub_item_input_price"`
	SubItemLimit      int     `json:"sub_item_limit"`
}

func (g AddAddOnDealSubItemSubItemListRequestEntity) String() string {
	return lib.ObjectToString(g)
}
