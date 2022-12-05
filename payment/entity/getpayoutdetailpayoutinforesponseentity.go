package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetPayoutDetailPayoutInfoResponseEntity struct {
	FromCurrency   string  `json:"from_currency"`
	PayoutCurrency string  `json:"payout_currency"`
	FromAmount     float32 `json:"from_amount"`
	PayoutAmount   float32 `json:"payout_amount"`
	ExchangeRate   string  `json:"exchange_rate"`
	PayoutTime     int     `json:"payout_time"`
	PayService     string  `json:"pay_service"`
	PayeeId        string  `json:"payee_id"`
}

func (g GetPayoutDetailPayoutInfoResponseEntity) String() string {
	return lib.ObjectToString(g)
}
