package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetOfferToggleStatusRequest struct {
}

func (g GetOfferToggleStatusRequest) String() string {
	return lib.ObjectToString(g)
}
