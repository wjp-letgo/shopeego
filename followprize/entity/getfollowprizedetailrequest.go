package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetFollowPrizeDetailRequest struct {
	CampaignId int64 `json:"campaign_id"`
}

func (g GetFollowPrizeDetailRequest) String() string {
	return lib.ObjectToString(g)
}
