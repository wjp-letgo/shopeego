package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteTopPicksResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	DeleteTopPicksResponseResponseEntity	`json:"response"`
}
func (g DeleteTopPicksResult) String() string {
    return lib.ObjectToString(g)
}
