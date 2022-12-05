package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetShopsByPartnerRequest struct {
	PageSize int `json:"page_size"`
	PageNo   int `json:"page_no"`
}

func (g GetShopsByPartnerRequest) String() string {
	return lib.ObjectToString(g)
}
