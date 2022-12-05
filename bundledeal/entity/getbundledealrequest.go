package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetBundleDealRequest struct {
	BundleDealId int64 `json:"bundle_deal_id"`
}

func (g GetBundleDealRequest) String() string {
	return lib.ObjectToString(g)
}
