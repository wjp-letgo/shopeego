package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetChannelListRequest struct {
	Region string `json:"region"`
}

func (g GetChannelListRequest) String() string {
	return lib.ObjectToString(g)
}
