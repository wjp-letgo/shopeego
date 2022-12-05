package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type AddBundleDealItemResponseResponseEntity struct {
	FailedList  []AddBundleDealItemFailedListResponseEntity `json:"failed_list"`
	SuccessList []int64                                     `json:"success_list"`
}

func (g AddBundleDealItemResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
