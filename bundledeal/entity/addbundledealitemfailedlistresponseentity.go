package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type AddBundleDealItemFailedListResponseEntity struct {
	ItemId      int64  `json:"item_id"`
	FailError   string `json:"fail_error"`
	FailMessage string `json:"fail_message"`
}

func (g AddBundleDealItemFailedListResponseEntity) String() string {
	return lib.ObjectToString(g)
}
