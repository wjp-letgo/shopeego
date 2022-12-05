package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type ShopPenaltyRequest struct {
}

func (g ShopPenaltyRequest) String() string {
	return lib.ObjectToString(g)
}
