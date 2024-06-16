package config

import (
	"github.com/shopspring/decimal"
)

type Config struct {
	DSN        string          `json:"dsn"`
	FloatRaio  decimal.Decimal `json:"float_ratio"`
	TradeRatio decimal.Decimal `json:"trade_ratio"`
	MaxValue   decimal.Decimal `json:"max_value"`
	MinValue   decimal.Decimal `json:"min_value"`
}

func init() {
}
