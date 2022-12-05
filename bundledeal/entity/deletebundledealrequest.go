package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteBundleDealRequest struct {
	BundleDealId int64 `json:"bundle_deal_id"`
}

func (g DeleteBundleDealRequest) String() string {
	return lib.ObjectToString(g)
}
