package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetMerchantsByPartnerRequest struct {
	PageSize int `json:"page_size"`
	PageNo   int `json:"page_no"`
}

func (g GetMerchantsByPartnerRequest) String() string {
	return lib.ObjectToString(g)
}
