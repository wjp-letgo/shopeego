package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetShopInstallmentStatusRequest struct {
}

func (g GetShopInstallmentStatusRequest) String() string {
	return lib.ObjectToString(g)
}
