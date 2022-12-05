package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetDiscountListRequest struct {
	DiscountStatus string `json:"discount_status"`
	PageNo         int    `json:"page_no"`
	PageSize       int    `json:"page_size"`
}

func (g GetDiscountListRequest) String() string {
	return lib.ObjectToString(g)
}
