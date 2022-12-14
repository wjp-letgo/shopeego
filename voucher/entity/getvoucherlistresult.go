package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetVoucherListResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	GetVoucherListResponseResponseEntity	`json:"response"`
}
func (g GetVoucherListResult) String() string {
    return lib.ObjectToString(g)
}
