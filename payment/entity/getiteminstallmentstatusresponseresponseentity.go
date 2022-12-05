package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetItemInstallmentStatusResponseResponseEntity struct {
	ItemInstallmentList []GetItemInstallmentStatusItemInstallmentListResponseEntity `json:"item_installment_list"`
}

func (g GetItemInstallmentStatusResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
