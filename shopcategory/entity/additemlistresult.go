package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type AddItemListResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	AddItemListResponseResponseEntity	`json:"response"`
}
func (g AddItemListResult) String() string {
    return lib.ObjectToString(g)
}
