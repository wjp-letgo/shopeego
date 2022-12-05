package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//DeleteItemResult
type DeleteItemResult struct {
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func (g DeleteItemResult) String() string {
	return lib.ObjectToString(g)
}
