package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetDetailRequest struct {
	FirstMileTrackingNumber string `json:"first_mile_tracking_number"`
	Cursor                  string `json:"cursor"`
}

func (g GetDetailRequest) String() string {
	return lib.ObjectToString(g)
}
