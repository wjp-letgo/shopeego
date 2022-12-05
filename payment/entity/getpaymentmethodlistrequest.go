package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetPaymentMethodListRequest struct {
}

func (g GetPaymentMethodListRequest) String() string {
	return lib.ObjectToString(g)
}
