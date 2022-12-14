package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateBundleDealResponseResponseEntity struct {
	BundleDealId   int64                                        `json:"bundle_deal_id"`
	Name           string                                       `json:"name"`
	StartTime      int                                          `json:"start_time"`
	EndTime        int                                          `json:"end_time"`
	BundleDealRule UpdateBundleDealBundleDealRuleResponseEntity `json:"bundle_deal_rule"`
	PurchaseLimit  int                                          `json:"purchase_limit"`
}

func (g UpdateBundleDealResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
