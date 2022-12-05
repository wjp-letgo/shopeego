package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateTopPicksResponseResponseEntity struct{
    CollectionList	[]UpdateTopPicksCollectionListResponseEntity	`json:"collection_list"`
}
func (g UpdateTopPicksResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
