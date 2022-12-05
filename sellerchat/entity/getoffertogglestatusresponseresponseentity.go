package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetOfferToggleStatusResponseResponseEntity struct {
	ShopId          int64  `json:"shop_id"`
	MakeOfferStatus string `json:"make_offer_status"`
}

func (g GetOfferToggleStatusResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
