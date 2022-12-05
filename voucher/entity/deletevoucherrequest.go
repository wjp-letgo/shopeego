package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteVoucherRequest struct{
    VoucherId	int64	`json:"voucher_id"`
}
func (g DeleteVoucherRequest) String() string {
    return lib.ObjectToString(g)
}
