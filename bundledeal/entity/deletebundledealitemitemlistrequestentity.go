package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteBundleDealItemItemListRequestEntity struct {
	ItemId int64 `json:"item_id"`
}

func (g DeleteBundleDealItemItemListRequestEntity) String() string {
	return lib.ObjectToString(g)
}
