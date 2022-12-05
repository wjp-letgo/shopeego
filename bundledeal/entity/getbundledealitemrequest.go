package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetBundleDealItemRequest struct {
	BundleDealId int64 `json:"bundle_deal_id"`
}

func (g GetBundleDealItemRequest) String() string {
	return lib.ObjectToString(g)
}
