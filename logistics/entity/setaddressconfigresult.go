package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//AddressTypeConfigEntity
type SetAddressConfigResult struct {
	commonentity.Result
}

//String
func (a SetAddressConfigResult) String() string {
	return lib.ObjectToString(a)
}
