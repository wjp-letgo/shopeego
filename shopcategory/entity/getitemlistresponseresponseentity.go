package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetItemListResponseResponseEntity struct{
    ItemList	[]int	`json:"item_list"`
}
func (g GetItemListResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
