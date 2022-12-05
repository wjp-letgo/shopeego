package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetReturnDetailUserResponseEntity struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Portrait string `json:"portrait"`
}

func (g GetReturnDetailUserResponseEntity) String() string {
	return lib.ObjectToString(g)
}
