package types

import (
	"github.com/shopspring/decimal"
)

type Asset struct {
	AvailBal      decimal.Decimal `json:"availBal"`  // available balance
	CashBal       decimal.Decimal `json:"cashBal"`   // total balance
	Ccy           string          `json:"ccy"`       // concurrency
	EqUsd         decimal.Decimal `json:"eqUsd"`     // USD value
	FixedBal      decimal.Decimal `json:"fixedBal"`  // fixed balance (?)
	FrozenBal     decimal.Decimal `json:"frozenBal"` // frozen balance
	AvailEq       decimal.Decimal `json:"availEq"`
	BorrowFroz    string          `json:"borrowFroz"`
	CrossLiab     string          `json:"crossLiab"`
	DisEq         string          `json:"disEq"`
	Eq            string          `json:"eq"`
	Imr           string          `json:"imr"`
	Interest      string          `json:"interest"`
	IsoEq         string          `json:"isoEq"`
	IsoLiab       string          `json:"isoLiab"`
	IsoUpl        string          `json:"isoUpl"`
	Liab          string          `json:"liab"`
	MaxLoan       string          `json:"maxLoan"`
	MgnRatio      string          `json:"mgnRatio"`
	Mmr           string          `json:"mmr"`
	NotionalLever string          `json:"notionalLever"`
	OrdFrozen     string          `json:"ordFrozen"`
	RewardBal     string          `json:"rewardBal"`
	SmtSyncEq     string          `json:"smtSyncEq"`
	SpotInUseAmt  string          `json:"spotInUseAmt"`
	SpotIsoBal    string          `json:"spotIsoBal"`
	StgyEq        string          `json:"stgyEq"`
	Twap          string          `json:"twap"`
	UTime         string          `json:"uTime"`
	Upl           string          `json:"upl"`
	UplLiab       string          `json:"uplLiab"`
}

type Balance struct {
	AdjEq       string          `json:"adjEq"`
	BorrowFroz  string          `json:"borrowFroz"`
	Details     []Asset         `json:"details"`
	Imr         string          `json:"imr"`
	IsoEq       string          `json:"isoEq"`
	MgnRatio    string          `json:"mgnRatio"`
	Mmr         string          `json:"mmr"`
	NotionalUsd string          `json:"notionalUsd"`
	OrdFroz     string          `json:"ordFroz"`
	TotalEq     decimal.Decimal `json:"totalEq"`
	UTime       string          `json:"uTime"`
	Upl         string          `json:"upl"`
}

type BalanceResponseV5 struct {
	Code string    `json:"code"`
	Msg  string    `json:"msg"`
	Data []Balance `json:"data"`
}
