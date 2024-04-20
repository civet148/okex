package types

import "github.com/shopspring/decimal"

const (
	TradeBaseUSDT = "USDT"
)

type TradeSide string

const (
	TradeSide_Buy  TradeSide = "buy"
	TradeSide_Sell TradeSide = "sell"
)

type OrderType string

const (
	OrderType_Limit  OrderType = "limit"  //限价单
	OrderType_Market OrderType = "market" //市价单
)

type TradeMode string

const (
	TradeMode_Cash TradeMode = "cash" //现金
)

type TradeRequest struct {
	Side      TradeSide       //* 买/卖
	OrderType OrderType       //* 订单类型
	TradeMode TradeMode       //* 交易模式
	Ccy       string          //* 代币
	Base      string          //* 币币单位(例如: USDT)
	Price     decimal.Decimal //* 价格
	Quantity  decimal.Decimal //* 数量
	OrderNo   string          //o 自定义订单号(选填)
}

type TradeResult struct {
	ClOrdId string `json:"clOrdId"`
	OrdId   string `json:"ordId"`
	Tag     string `json:"tag"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}

type TradeResponseV5 struct {
	Code    string        `json:"code"`
	Msg     string        `json:"msg"`
	Data    []TradeResult `json:"data"`
	InTime  string        `json:"inTime"`
	OutTime string        `json:"outTime"`
}
