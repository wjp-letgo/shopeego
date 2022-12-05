package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type ShopPerformanceLateShipmentRateResponseEntity struct {
	TotalData ShopPerformanceTotalDataResponseEntity `json:"total_data"`
}

func (g ShopPerformanceLateShipmentRateResponseEntity) String() string {
	return lib.ObjectToString(g)
}
