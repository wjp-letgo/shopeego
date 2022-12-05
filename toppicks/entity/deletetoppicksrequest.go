package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteTopPicksRequest struct{
    TopPicksId	int64	`json:"top_picks_id"`
}
func (g DeleteTopPicksRequest) String() string {
    return lib.ObjectToString(g)
}
