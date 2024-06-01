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
	OrderType_Fok    OrderType = "fok"    //限价单(全部卖出或取消)
	OrderType_Ioc    OrderType = "ioc"    //限价单(立刻成交并取消剩余)
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

type TradeOrder struct {
	AccFillSz         string          `json:"accFillSz"`
	AvgPx             string          `json:"avgPx"`
	CTime             string          `json:"cTime"`
	Category          string          `json:"category"`
	Ccy               string          `json:"ccy"`
	ClOrdId           string          `json:"clOrdId"`
	Fee               string          `json:"fee"`
	FeeCcy            string          `json:"feeCcy"`
	FillPx            string          `json:"fillPx"`
	FillSz            string          `json:"fillSz"`
	FillTime          string          `json:"fillTime"`
	InstId            string          `json:"instId"`
	InstType          string          `json:"instType"`
	Lever             string          `json:"lever"`
	OrdId             string          `json:"ordId"`
	OrdType           string          `json:"ordType"`
	Pnl               string          `json:"pnl"`
	PosSide           string          `json:"posSide"`
	Px                decimal.Decimal `json:"px"`
	PxUsd             string          `json:"pxUsd"`
	PxVol             string          `json:"pxVol"`
	PxType            string          `json:"pxType"`
	Rebate            string          `json:"rebate"`
	RebateCcy         string          `json:"rebateCcy"`
	Side              string          `json:"side"`
	AttachAlgoClOrdId string          `json:"attachAlgoClOrdId"`
	SlOrdPx           string          `json:"slOrdPx"`
	SlTriggerPx       string          `json:"slTriggerPx"`
	SlTriggerPxType   string          `json:"slTriggerPxType"`
	AttachAlgoOrds    []interface{}   `json:"attachAlgoOrds"`
	Source            string          `json:"source"`
	State             string          `json:"state"`
	StpId             string          `json:"stpId"`
	StpMode           string          `json:"stpMode"`
	Sz                decimal.Decimal `json:"sz"`
	Tag               string          `json:"tag"`
	TdMode            string          `json:"tdMode"`
	TgtCcy            string          `json:"tgtCcy"`
	TpOrdPx           string          `json:"tpOrdPx"`
	TpTriggerPx       string          `json:"tpTriggerPx"`
	TpTriggerPxType   string          `json:"tpTriggerPxType"`
	TradeId           string          `json:"tradeId"`
	ReduceOnly        string          `json:"reduceOnly"`
	QuickMgnType      string          `json:"quickMgnType"`
	AlgoClOrdId       string          `json:"algoClOrdId"`
	AlgoId            string          `json:"algoId"`
	IsTpLimit         string          `json:"isTpLimit"`
	UTime             string          `json:"uTime"`
}

type PendingOrdersResponseV5 struct {
	Code string       `json:"code"`
	Msg  string       `json:"msg"`
	Data []TradeOrder `json:"data"`
}
