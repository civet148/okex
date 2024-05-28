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

func NewOkexClient(apiKey *types.APIKeyInfo, strUrl string, isDebug bool) *OkexClient {
	client := rest.NewRESTClient(strUrl, apiKey, false, isDebug)
	return &OkexClient{
		client: client,
	}
}

func (m *OkexClient) Balance() (response *types.BalanceResponseV5, err error) {
	var res *rest.RESTAPIResult
	res, err = m.client.Get(context.Background(), types.API_V5_ACCOUNT_BALANCE, nil)
	if err != nil {
		return nil, log.Errorf("GET balance error [%s]", err.Error())
	}
	response = &types.BalanceResponseV5{}
	err = json.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		return nil, log.Errorf("response body json unmarshal error [%s]", err.Error())
	}
	if response.Code != "0" {
		return nil, log.Errorf("error code [%v] message [%s]", response.Code, response.Msg)
	}
	return response, nil
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
	res, err = m.client.Post(context.Background(), types.API_V5_TRADE_ORDER, &params)
	if err != nil {
		return "", log.Errorf("POST trade order error [%s]", err.Error())
	}
	log.Debugf("%s", res.Body)
	response := &types.TradeResponseV5{}
	err = json.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		return "", log.Errorf("response body json unmarshal error [%s]", err.Error())
	}
	if response.Code != "0" {
		return "", log.Errorf("error code [%v] message [%s]", response.Code, response.Msg)
	}
	orderId = response.Data[0].OrdId
	return
}

// SpotPendingOrders 列出所有现货挂单
// GET /api/v5/trade/orders-pending?ordType=post_only,fok,ioc&instType=SPOT
// 参考文档：https://www.okx.com/docs-v5/zh/#order-book-trading-trade-get-order-list
func (m *OkexClient) SpotPendingOrders(strOrderType string, instIds ...string) (orders []types.TradeOrder, err error) {
	var params = map[string]interface{}{
		"instType": "SPOT",
	}
	if strOrderType != "" {
		params["orderType"] = strOrderType
	}
	var res *rest.RESTAPIResult
	res, err = m.client.Get(context.Background(), types.API_V5_PENDING_ORDERS, &params)
	if err != nil {
		return nil, log.Errorf(err.Error())
	}

	log.Debugf("%s", res.Body)
	response := &types.PendingOrdersResponseV5{}
	err = json.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		return nil, log.Errorf("response body json unmarshal error [%s]", err.Error())
	}
	if response.Code != "0" {
		return nil, log.Errorf("error code [%v] message [%s]", response.Code, response.Msg)
	}
	log.Debugf("order count [%d]", len(response.Data))
	return response.Data, nil
}

func (m *OkexClient) SpotCancelOrder(instId, strOrderId string) (err error) {
	var params = map[string]interface{}{
		"instId": instId,
		"ordId":  strOrderId,
	}
	var res *rest.RESTAPIResult
	res, err = m.client.Post(context.Background(), types.API_V5_CANCEL_ORDER, &params)
	if err != nil {
		return log.Errorf("POST cancel order error [%s]", err.Error())
	}
	log.Debugf("%s", res.Body)
	response := &types.CancelOrderResponseV5{}
	err = json.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		return log.Errorf("response body json unmarshal error [%s]", err.Error())
	}
	if response.Code != "0" {
		return log.Errorf("error code [%v] message [%s]", response.Code, response.Msg)
	}
	return nil
}

func (m *OkexClient) SpotPrice(instId string) (price types.InstPrice, err error) {
	var params = map[string]interface{}{
		"instId": instId,
	}
	var res *rest.RESTAPIResult
	res, err = m.client.Get(context.Background(), types.API_V5_MARKET_TICKER, &params)
	if err != nil {
		return price, log.Errorf("POST cancel order error [%s]", err.Error())
	}
	log.Debugf("%s", res.Body)
	response := &types.MarketTickerResponseV5{}
	err = json.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		return price, log.Errorf("response body json unmarshal error [%s]", err.Error())
	}
	if response.Code != "0" {
		return price, log.Errorf("error code [%v] message [%s]", response.Code, response.Msg)
	}
	for _, v := range response.Data {
		if v.InstType == types.SPOT {
			return v, nil
		}
	}
	return price, nil
}
