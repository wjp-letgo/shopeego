package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetAddOnDealListRequest struct {
	PromotionStatus string `json:"promotion_status"`
	PageNo          int    `json:"page_no"`
	PageSize        int    `json:"page_size"`
}

func (g GetAddOnDealListRequest) String() string {
	return lib.ObjectToString(g)
}
