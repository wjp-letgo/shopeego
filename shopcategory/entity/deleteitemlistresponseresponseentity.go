package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DeleteItemListResponseResponseEntity struct{
    ShopCategoryId	int64	`json:"shop_category_id"`
    InvalidItemId	int64	`json:"invalid_item_id"`
    CurrentCount	int	`json:"current_count"`
}
func (g DeleteItemListResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
