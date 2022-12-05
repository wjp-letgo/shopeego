package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type SetItemInstallmentStatusRequest struct {
	ItemIdList []int64 `json:"item_id_list"`
	TenureList []int   `json:"tenure_list"`
}

func (g SetItemInstallmentStatusRequest) String() string {
	return lib.ObjectToString(g)
}
