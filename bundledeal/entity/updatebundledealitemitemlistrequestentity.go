package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateBundleDealItemItemListRequestEntity struct {
	ItemId int64 `json:"item_id"`
	Status int   `json:"status"`
}

func (g UpdateBundleDealItemItemListRequestEntity) String() string {
	return lib.ObjectToString(g)
}
