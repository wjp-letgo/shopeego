package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type SetItemInstallmentStatusResult struct {
	Error     string                                         `json:"error"`
	Message   string                                         `json:"message"`
	RequestId string                                         `json:"request_id"`
	Response  SetItemInstallmentStatusResponseResponseEntity `json:"response"`
}

func (g SetItemInstallmentStatusResult) String() string {
	return lib.ObjectToString(g)
}
