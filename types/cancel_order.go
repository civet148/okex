package types

type OrderResult struct {
	ClOrdId string `json:"clOrdId"`
	OrdId   string `json:"ordId"`
	Ts      string `json:"ts"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}

type CancelOrderResponseV5 struct {
	Code    string        `json:"code"`
	Msg     string        `json:"msg"`
	Data    []OrderResult `json:"data"`
	InTime  string        `json:"inTime"`
	OutTime string        `json:"outTime"`
}
