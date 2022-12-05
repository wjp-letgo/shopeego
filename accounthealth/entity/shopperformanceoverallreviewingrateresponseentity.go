package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type ShopPerformanceOverallReviewingRateResponseEntity struct {
	TotalData ShopPerformanceTotalDataResponseEntity `json:"total_data"`
}

func (g ShopPerformanceOverallReviewingRateResponseEntity) String() string {
	return lib.ObjectToString(g)
}
