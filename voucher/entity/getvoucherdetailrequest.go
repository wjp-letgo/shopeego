package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetVoucherDetailRequest struct{
    VoucherId	int64	`json:"voucher_id"`
}
func (g GetVoucherDetailRequest) String() string {
    return lib.ObjectToString(g)
}
