package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//UpdateGlobalModelResult
type UpdateGlobalModelResult struct {
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func (g UpdateGlobalModelResult) String() string {
	return lib.ObjectToString(g)
}
