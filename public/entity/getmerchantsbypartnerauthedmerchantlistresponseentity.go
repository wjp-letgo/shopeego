package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetMerchantsByPartnerAuthedMerchantListResponseEntity struct {
	Region     string `json:"region"`
	MerchantId int64  `json:"merchant_id"`
	AuthTime   int    `json:"auth_time"`
	ExpireTime int    `json:"expire_time"`
}

func (g GetMerchantsByPartnerAuthedMerchantListResponseEntity) String() string {
	return lib.ObjectToString(g)
}
