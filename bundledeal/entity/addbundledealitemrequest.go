package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type AddBundleDealItemRequest struct {
	BundleDealId int64                                    `json:"bundle_deal_id"`
	ItemList     []AddBundleDealItemItemListRequestEntity `json:"item_list"`
}

func (g AddBundleDealItemRequest) String() string {
	return lib.ObjectToString(g)
}
