package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UpdateShopCategoryResponseResponseEntity struct{
    ShopCategoryId	int64	`json:"shop_category_id"`
    Name	string	`json:"name"`
    SortWeight	int	`json:"sort_weight"`
    Status	string	`json:"status"`
}
func (g UpdateShopCategoryResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
