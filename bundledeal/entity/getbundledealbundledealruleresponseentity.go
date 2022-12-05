package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetBundleDealBundleDealRuleResponseEntity struct {
	RuleType           int     `json:"rule_type"`
	DiscountValue      float32 `json:"discount_value"`
	FixPrice           float32 `json:"fix_price"`
	DiscountPercentage int     `json:"discount_percentage"`
	MinAmount          int     `json:"min_amount"`
}

func (g GetBundleDealBundleDealRuleResponseEntity) String() string {
	return lib.ObjectToString(g)
}
