package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type ShopPerformanceCustomerSatisfactionResponseEntity struct {
	OverallReviewingRate ShopPerformanceOverallReviewingRateResponseEntity `json:"overall_reviewing_rate"`
}

func (g ShopPerformanceCustomerSatisfactionResponseEntity) String() string {
	return lib.ObjectToString(g)
}
