package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type SetOfferToggleStatusRequest struct {
	MakeOfferStatus string `json:"make_offer_status	"`
}

func (g SetOfferToggleStatusRequest) String() string {
	return lib.ObjectToString(g)
}
