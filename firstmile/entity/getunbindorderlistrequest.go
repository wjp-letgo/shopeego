package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetUnbindOrderListRequest struct {
	Cursor                 string `json:"cursor"`
	PageSize               int    `json:"page_size"`
	ResponseOptionalFields string `json:"response_optional_fields"`
}

func (g GetUnbindOrderListRequest) String() string {
	return lib.ObjectToString(g)
}
