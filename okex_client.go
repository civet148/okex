package okex

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/civet148/log"
	"github.com/civet148/okex/rest"
	"github.com/civet148/okex/types"
	"strings"
)

type OkexClient struct {
	client *rest.RESTAPI
}

func NewOkexClient(apiKey *types.APIKeyInfo, strUrl string) *OkexClient {
	client := rest.NewRESTClient(strUrl, apiKey, false)
	return &OkexClient{
		client: client,
	}
}

func (m *OkexClient) Balance() (balance *types.BalanceResponseV5, err error) {
	var res *rest.RESTAPIResult
	res, err = m.client.Get(context.Background(), types.ApiV5AccountBalance, nil)
	if err != nil {
		return nil, log.Errorf("GET balance error [%s]", err.Error())
	}
	balance = &types.BalanceResponseV5{}
	err = json.Unmarshal([]byte(res.Body), &balance)
	if err != nil {
		return nil, log.Errorf("response body json unmarshal error [%s]", err.Error())
	}
	if balance.Code != "0" {
		return nil, log.Errorf("%s", balance.Msg)
	}
	return balance, nil
}

/*
---------------------------- REQUEST ----------------------------

	{
	    "instId":"BTC-USDT",
	    "tdMode":"cash",
	    "clOrdId":"b15",
	    "side":"buy",
	    "ordType":"limit",
	    "px":"2.15",
	    "sz":"2"
	}

---------------------------- RESPONSE ----------------------------

	{
	    "code":"0",
	    "msg":"",
	    "data":[
	        {
	            "clOrdId":"oktswap6",
	            "ordId":"12345689",
	            "tag":"",
	            "sCode":"0",
	            "sMsg":""
	        }
	    ],
	    "inTime": "1695190491421339",
	    "outTime": "1695190491423240"
	}
*/
func (m *OkexClient) SpotTradeOrder(req *types.TradeRequest) (orderId string, err error) {
	var instId = fmt.Sprintf("%s-%s", strings.ToUpper(req.Ccy), strings.ToUpper(req.Base))
	var params = map[string]interface{}{
		"instId":  instId,
		"tdMode":  req.TradeMode,
		"clOrdId": req.OrderNo,
		"side":    req.Side,
		"ordType": req.OrderType,
		"px":      req.Price.String(),
		"sz":      req.Quantity.String(),
	}
	var res *rest.RESTAPIResult
	res, err = m.client.Post(context.Background(), types.ApiV5AccountTradeOrder, &params)
	if err != nil {
		return "", log.Errorf("POST trade order error [%s]", err.Error())
	}
	log.Infof("%s", res.Body)
	response := &types.TradeResponseV5{}
	err = json.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		return "", log.Errorf("response body json unmarshal error [%s]", err.Error())
	}
	if response.Code != "0" {
		return "", log.Errorf("%s", response.Msg)
	}
	if len(response.Data) == 0 {
		return "", log.Errorf("empty result")
	}
	orderId = response.Data[0].OrdId
	return
}
