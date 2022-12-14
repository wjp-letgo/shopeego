package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type SetShopInstallmentStatusRequest struct {
	InstallmentStatus int `json:"installment_status"`
}

func (g SetShopInstallmentStatusRequest) String() string {
	return lib.ObjectToString(g)
}
