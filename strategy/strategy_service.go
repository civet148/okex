package strategy

import "github.com/civet148/okex/config"

type StrategyService struct {
	polices config.OkexPolices
}

func NewStrategyService(polices config.OkexPolices) *StrategyService {
	return &StrategyService{
		polices: polices,
	}
}

func (m *StrategyService) Run() error {

	return nil
}
