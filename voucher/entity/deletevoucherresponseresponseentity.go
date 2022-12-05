package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteVoucherResponseResponseEntity struct{
    VoucherId	int64	`json:"voucher_id"`
}
func (g DeleteVoucherResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
