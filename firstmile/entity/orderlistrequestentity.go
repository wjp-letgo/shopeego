package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type OrderListRequestEntity struct {
	OrderSn       string `json:"order_sn"`
	PackageNumber string `json:"package_number"`
}

func (g OrderListRequestEntity) String() string {
	return lib.ObjectToString(g)
}
