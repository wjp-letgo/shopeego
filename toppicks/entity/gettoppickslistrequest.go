package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetTopPicksListRequest struct{
}
func (g GetTopPicksListRequest) String() string {
    return lib.ObjectToString(g)
}
