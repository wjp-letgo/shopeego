package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetPaymentMethodListResponseResponseEntity struct {
	PaymentMethod []string `json:"payment_method"`
	Region        string   `json:"region"`
}

func (g GetPaymentMethodListResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
