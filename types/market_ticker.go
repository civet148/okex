package types

import "github.com/shopspring/decimal"

type MarketPrice struct {
	InstType  string          `json:"instType"`  //产品类型(现货: SPOT)
	InstId    string          `json:"instId"`    //产品ID（例如：BTC-USDT）
	Last      decimal.Decimal `json:"last"`      //最新成交价
	LastSz    decimal.Decimal `json:"lastSz"`    //最新成交的数量
	AskPx     decimal.Decimal `json:"askPx"`     //卖一价
	AskSz     decimal.Decimal `json:"askSz"`     //卖一价对应的数量
	BidPx     decimal.Decimal `json:"bidPx"`     //买一价
	BidSz     decimal.Decimal `json:"bidSz"`     //买一价对应数量
	Open24H   decimal.Decimal `json:"open24h"`   //24小时开盘价
	High24H   decimal.Decimal `json:"high24h"`   //24小时最高价
	Low24H    decimal.Decimal `json:"low24h"`    //24小时最低价
	VolCcy24H decimal.Decimal `json:"volCcy24h"` //24小时开盘价(以币为单位)
	Vol24H    decimal.Decimal `json:"vol24h"`    //24小时开盘价(以张为单位)
	SodUtc0   string          `json:"sodUtc0"`   //UTC+0 时开盘价
	SodUtc8   string          `json:"sodUtc8"`   //UTC+8 时开盘价
	Ts        string          `json:"ts"`        //ticker数据产生时间，Unix时间戳的毫秒数格式
}

type MarketTickerResponseV5 struct {
	Code string      `json:"code"`
	Msg  string        `json:"msg"`
	Data []MarketPrice `json:"data"`
}
