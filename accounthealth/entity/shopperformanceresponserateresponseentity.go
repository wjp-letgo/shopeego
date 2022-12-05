package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type ShopPerformanceResponseRateResponseEntity struct {
	TotalData ShopPerformanceTotalDataResponseEntity `json:"total_data"`
}

func (g ShopPerformanceResponseRateResponseEntity) String() string {
	return lib.ObjectToString(g)
}
