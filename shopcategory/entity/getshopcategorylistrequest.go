package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetShopCategoryListRequest struct{
    PageSize	int	`json:"page_size"`
    PageNo	int	`json:"page_no"`
}
func (g GetShopCategoryListRequest) String() string {
    return lib.ObjectToString(g)
}
