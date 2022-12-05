package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type ShopPerformanceRequest struct {
}

func (g ShopPerformanceRequest) String() string {
	return lib.ObjectToString(g)
}
