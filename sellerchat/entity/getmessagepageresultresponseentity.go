package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetMessagePageResultResponseEntity struct {
	NextOffset string `json:"next_offset"`
	PageSize   int    `json:"page_size"`
}

func (g GetMessagePageResultResponseEntity) String() string {
	return lib.ObjectToString(g)
}
