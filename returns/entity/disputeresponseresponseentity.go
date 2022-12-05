package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DisputeResponseResponseEntity struct {
	ReturnSn string `json:"return_sn"`
}

func (g DisputeResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
