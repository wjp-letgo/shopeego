package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type AddBundleDealItemItemListRequestEntity struct {
	ItemId int64 `json:"item_id"`
	Status int   `json:"status"`
}

func (g AddBundleDealItemItemListRequestEntity) String() string {
	return lib.ObjectToString(g)
}
