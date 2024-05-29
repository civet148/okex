package types

type LoanToken struct {
	Ccy          string `json:"ccy"`
	InterestRate string `json:"interestRate"`
}

type LoanTokenResponseV5 struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data []LoanToken `json:"data"`
}
