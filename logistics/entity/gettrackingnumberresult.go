package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//GetTrackingNumberResult
type GetTrackingNumberResult struct {
	commonentity.Result
	Response TrackingNumberEntity `json:"response"`
}

//String
func (g GetTrackingNumberResult) String() string {
	return lib.ObjectToString(g)
}
