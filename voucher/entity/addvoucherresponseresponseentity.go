package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type AddVoucherResponseResponseEntity struct{
    VoucherId	int64	`json:"voucher_id"`
}
func (g AddVoucherResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
