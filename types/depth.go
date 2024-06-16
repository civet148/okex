package types

import "github.com/shopspring/decimal"

type DepthUnit struct {
	Price         decimal.Decimal
	Quantity      decimal.Decimal
	OrderQuantity decimal.Decimal
}

type OrderBook struct {
	Bids []DepthUnit
	Asks []DepthUnit
}

type OrderBookResponseV5 struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Asks [][]decimal.Decimal `json:"asks"`
		Bids [][]decimal.Decimal `json:"bids"`
		Ts   string              `json:"ts"`
	} `json:"data"`
}
