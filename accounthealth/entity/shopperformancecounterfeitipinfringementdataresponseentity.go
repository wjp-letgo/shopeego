package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type ShopPerformanceCounterfeitIpInfringementDataResponseEntity struct {
	Target            string `json:"target"`
	MyShopPerformance string `json:"my_shop_performance	"`
	PenaltyPoints     string `json:"penalty_points"`
}

func (g ShopPerformanceCounterfeitIpInfringementDataResponseEntity) String() string {
	return lib.ObjectToString(g)
}
