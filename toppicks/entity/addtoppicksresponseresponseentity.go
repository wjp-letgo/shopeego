package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type AddTopPicksResponseResponseEntity struct{
    CollectionList	[]AddTopPicksCollectionListResponseEntity	`json:"collection_list"`
}
func (g AddTopPicksResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
