package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteShopCategoryRequest struct{
    ShopCategoryId	int64	`json:"shop_category_id"`
}
func (g DeleteShopCategoryRequest) String() string {
    return lib.ObjectToString(g)
}
