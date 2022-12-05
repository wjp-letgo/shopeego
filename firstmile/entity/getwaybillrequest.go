package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetWaybillRequest struct {
	FirstMileTrackingNumberList []string `json:"first_mile_tracking_number_list"`
}

func (g GetWaybillRequest) String() string {
	return lib.ObjectToString(g)
}
