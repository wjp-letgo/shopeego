package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetItemListRequest struct{
    ShopCategoryId	int64	`json:"shop_category_id"`
    PageSize	int	`json:"page_size"`
    PageNo	int	`json:"page_no"`
}
func (g GetItemListRequest) String() string {
    return lib.ObjectToString(g)
}
