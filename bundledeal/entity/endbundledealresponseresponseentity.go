package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type EndBundleDealResponseResponseEntity struct {
	BundleDealId int64 `json:"bundle_deal_id"`
}

func (g EndBundleDealResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
