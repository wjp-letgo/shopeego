package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteAddOnDealResponseResponseEntity struct {
	AddOnDealId int64 `json:"add_on_deal_id"`
}

func (g DeleteAddOnDealResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
