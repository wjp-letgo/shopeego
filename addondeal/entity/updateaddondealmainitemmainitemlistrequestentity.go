package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateAddOnDealMainItemMainItemListRequestEntity struct {
	ItemId int64 `json:"item_id"`
	Status int   `json:"status"`
}

func (g UpdateAddOnDealMainItemMainItemListRequestEntity) String() string {
	return lib.ObjectToString(g)
}
