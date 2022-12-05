package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GenerateFirstMileTrackingNumberSellerInfoRequestEntity struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Zipcode string `json:"zipcode"`
	Region  string `json:"region"`
	Phone   string `json:"phone"`
}

func (g GenerateFirstMileTrackingNumberSellerInfoRequestEntity) String() string {
	return lib.ObjectToString(g)
}
