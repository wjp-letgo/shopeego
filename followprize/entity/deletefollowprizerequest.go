package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteFollowPrizeRequest struct {
	CampaignId int64 `json:"campaign_id"`
}

func (g DeleteFollowPrizeRequest) String() string {
	return lib.ObjectToString(g)
}
