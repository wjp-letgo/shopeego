package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetReturnListRequest struct {
	PageNo         int `json:"page_no"`
	PageSize       int `json:"page_size"`
	CreateTimeFrom int `json:"create_time_from"`
	CreateTimeTo   int `json:"create_time_to"`
}

func (g GetReturnListRequest) String() string {
	return lib.ObjectToString(g)
}
