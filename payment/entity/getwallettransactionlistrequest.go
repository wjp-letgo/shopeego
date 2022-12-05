package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetWalletTransactionListRequest struct {
	PageNo          int    `json:"page_no"`
	PageSize        int    `json:"page_size"`
	CreateTimeFrom  int    `json:"create_time_from"`
	CreateTimeTo    int    `json:"create_time_to"`
	WalletType      string `json:"wallet_type"`
	TransactionType string `json:"transaction_type"`
}

func (g GetWalletTransactionListRequest) String() string {
	return lib.ObjectToString(g)
}
