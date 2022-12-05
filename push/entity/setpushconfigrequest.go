package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type SetPushConfigRequest struct {
	CallbackUrl   string                               `json:"callback_url"`
	PushConfig    SetPushConfigPushConfigRequestEntity `json:"push_config"`
	BlockedShopId []int64                              `json:"blocked_shop_id"`
}

func (g SetPushConfigRequest) String() string {
	return lib.ObjectToString(g)
}
