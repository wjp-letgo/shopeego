package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//AddGlobalModelResult
type AddGlobalModelResult struct {
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func (g AddGlobalModelResult) String() string {
	return lib.ObjectToString(g)
}
