package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetVoucherListRequest struct{
    PageNo	int	`json:"page_no"`
    PageSize	int	`json:"page_size"`
    Status	string	`json:"status"`
}
func (g GetVoucherListRequest) String() string {
    return lib.ObjectToString(g)
}
