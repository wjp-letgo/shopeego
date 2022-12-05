package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type AddBundleDealResponseResponseEntity struct {
	BundleDealId int64 `json:"bundle_deal_id"`
}

func (g AddBundleDealResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
