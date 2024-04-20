package config

import (
	_ "embed"
	"encoding/json"
	"github.com/shopspring/decimal"
	"os"
)

//go:embed config.json
var ConfigTemplateBytes []byte

type OkexPolices []OkexPolicy

type OkexPolicy struct {
	Ccy  string       `json:"ccy"`
	Buy  PolicyFields `json:"buy"`
	Sell PolicyFields `json:"sell"`
}

type PolicyFields struct {
	FluctuationRate decimal.Decimal `json:"fluctuation_rate"`
	TradeRate       decimal.Decimal `json:"trade_rate"`
	Max             decimal.Decimal `json:"max"`
	Min             decimal.Decimal `json:"min"`
}

func LoadConfig(strPath string) (polices []OkexPolicy, err error) {
	var data []byte
	data, err = os.ReadFile(strPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &polices)
	if err != nil {
		return nil, err
	}
	return
}

func GenerateConfig(strPath string) (err error) {
	err = os.WriteFile(strPath, ConfigTemplateBytes, os.ModePerm)
	return err
}
