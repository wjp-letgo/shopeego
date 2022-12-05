package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetPushConfigRequest struct {
}

func (g GetPushConfigRequest) String() string {
	return lib.ObjectToString(g)
}
