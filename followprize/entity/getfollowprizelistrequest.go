package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetFollowPrizeListRequest struct {
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
	Status   string `json:"status"`
}

func (g GetFollowPrizeListRequest) String() string {
	return lib.ObjectToString(g)
}
