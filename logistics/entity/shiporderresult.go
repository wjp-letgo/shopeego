package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//ShipOrderResult
type ShipOrderResult struct {
	commonentity.Result
}

//String
func (s ShipOrderResult) String() string {
	return lib.ObjectToString(s)
}
