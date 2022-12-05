package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateShopCategoryRequest struct{
    ShopCategoryId	int	`json:"shop_category_id"`
    Name	string	`json:"name"`
    SortWeight	int	`json:"sort_weight"`
    Status	string	`json:"status"`
}
func (g UpdateShopCategoryRequest) String() string {
    return lib.ObjectToString(g)
}
