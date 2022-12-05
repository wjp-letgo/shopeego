package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetUnbindOrderListResponseResponseEntity struct {
	More       bool                                        `json:"more"`
	OrderList  []GetUnbindOrderListOrderListResponseEntity `json:"order_list"`
	NextCursor string                                      `json:"next_cursor"`
}

func (g GetUnbindOrderListResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
