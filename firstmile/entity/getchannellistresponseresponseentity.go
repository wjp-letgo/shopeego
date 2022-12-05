package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetChannelListResponseResponseEntity struct {
	LogisticsChannelList []GetChannelListLogisticsChannelListResponseEntity `json:"logistics_channel_list"`
}

func (g GetChannelListResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
