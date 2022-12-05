package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type ConfirmResponseResponseEntity struct {
	ReturnSn string `json:"return_sn"`
}

func (g ConfirmResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
