package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type BindFirstMileTrackingNumberOrderListRequestEntity struct {
	OrderSn       string `json:"order_sn"`
	PackageNumber string `json:"package_number"`
}

func (g BindFirstMileTrackingNumberOrderListRequestEntity) String() string {
	return lib.ObjectToString(g)
}
