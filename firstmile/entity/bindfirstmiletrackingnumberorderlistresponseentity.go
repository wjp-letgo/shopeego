package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type BindFirstMileTrackingNumberOrderListResponseEntity struct {
	OrderSn       string `json:"order_sn"`
	PackageNumber string `json:"package_number"`
	FailError     string `json:"fail_error"`
	FailMessage   string `json:"fail_message"`
}

func (g BindFirstMileTrackingNumberOrderListResponseEntity) String() string {
	return lib.ObjectToString(g)
}
