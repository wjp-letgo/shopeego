package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetAddOnDealMainItemMainItemListResponseEntity struct {
	ItemId int64 `json:"item_id"`
	Status int   `json:"status"`
}

func (g GetAddOnDealMainItemMainItemListResponseEntity) String() string {
	return lib.ObjectToString(g)
}
