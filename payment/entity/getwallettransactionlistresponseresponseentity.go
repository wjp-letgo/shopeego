package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetWalletTransactionListResponseResponseEntity struct {
	TransactionList []GetWalletTransactionListTransactionListResponseEntity `json:"transaction_list"`
	More            bool                                                    `json:"more"`
}

func (g GetWalletTransactionListResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
