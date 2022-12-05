package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetShopsByPartnerSipAffiShopListResponseEntity struct {
	Region     string `json:"region"`
	AffiShopId int64  `json:"affi_shop_id"`
}

func (g GetShopsByPartnerSipAffiShopListResponseEntity) String() string {
	return lib.ObjectToString(g)
}
