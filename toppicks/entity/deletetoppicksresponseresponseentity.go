package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteTopPicksResponseResponseEntity struct{
    TopPicksId	int64	`json:"top_picks_id"`
}
func (g DeleteTopPicksResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
