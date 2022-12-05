package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type AddTopPicksRequest struct{
    Name	string	`json:"name"`
    ItemIdList	[]int64	`json:"item_id_list"`
    IsActivated	bool	`json:"is_activated"`
}
func (g AddTopPicksRequest) String() string {
    return lib.ObjectToString(g)
}
