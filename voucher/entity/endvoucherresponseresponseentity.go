package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type EndVoucherResponseResponseEntity struct{
    VoucherId	int64	`json:"voucher_id"`
}
func (g EndVoucherResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
