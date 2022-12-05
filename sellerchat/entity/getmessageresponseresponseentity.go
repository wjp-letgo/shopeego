package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetMessageResponseResponseEntity struct {
	Messages   []GetMessageMessagesResponseEntity `json:"messages"`
	PageResult GetMessagePageResultResponseEntity `json:"page_result"`
}

func (g GetMessageResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
