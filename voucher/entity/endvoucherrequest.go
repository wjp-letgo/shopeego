package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type EndVoucherRequest struct{
    VoucherId	int64	`json:"voucher_id"`
}
func (g EndVoucherRequest) String() string {
    return lib.ObjectToString(g)
}
