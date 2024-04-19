package okex

import (
	"context"
	_ "embed"
	"encoding/json"
	"github.com/civet148/log"
	"github.com/civet148/okex/rest"
	"github.com/civet148/okex/types"
)

type OkexClient struct {
	client *rest.RESTAPI
}

//go:embed config.json
var ConfigTemplateBytes []byte

func NewOkexClient(apiKey *types.APIKeyInfo, strUrl string) *OkexClient {
	client := rest.NewRESTClient(strUrl, apiKey, false)
	return &OkexClient{
		client: client,
	}
}

func (c *OkexClient) Balance() (balance *types.BalanceV5, err error) {
	res, err := c.client.Get(context.Background(), "/api/v5/account/balance", nil)
	if err != nil {
		return nil, log.Errorf("GET balance error [%s]", err.Error())
	}
	balance = &types.BalanceV5{}
	err = json.Unmarshal([]byte(res.Body), &balance)
	if err != nil {
		return nil, log.Errorf("response body json unmarshal error [%s]", err.Error())
	}
	return balance, nil
}
