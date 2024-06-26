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

type Options struct {
	ApiUrl         string
	TimeoutSeconds int
	IsDebug        bool
	IsSimulate     bool
}

func defaultOptions() *Options {
	return &Options{
		TimeoutSeconds: 30,
		IsDebug:        false,
		IsSimulate:     false,
	}
}
func NewOkexClient(apiKey *types.APIKeyInfo, options ...*Options) *OkexClient {
	var opts = defaultOptions()
	for _, o := range options {
		if o != nil {
			opts = o
		}
	}
	if opts.ApiUrl == "" {
		opts.ApiUrl = types.OKEX_API_ADDR
	}
	client := rest.NewRESTClient(opts.ApiUrl, apiKey, opts.IsSimulate, opts.IsDebug, opts.TimeoutSeconds)
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

func (m *OkexClient) SpotPrice(instId string) (price types.MarketPrice, err error) {
	var params = map[string]interface{}{
		"instId": instId,
	}
	var res *rest.RESTAPIResult
	res, err = m.client.Get(context.Background(), types.API_V5_MARKET_TICKER, &params)
	if err != nil {
		return price, log.Errorf(err.Error())
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

func (m *OkexClient) SpotPrices(instIds ...string) (prices []types.MarketPrice, err error) {
	var ids = make(map[string]bool)
	for _, instId := range instIds {
		ids[instId] = true
	}
	var res *rest.RESTAPIResult
	var params = map[string]interface{}{
		"instType": types.SPOT,
	}
	res, err = m.client.Get(context.Background(), types.API_V5_MARKET_TICKERS, &params)
	if err != nil {
		return nil, log.Errorf(err.Error())
	}
	log.Debugf("%s", res.Body)
	response := &types.MarketTickerResponseV5{}
	err = json.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		return nil, log.Errorf("response body [%s] json unmarshal error [%s]", res.Body, err.Error())
	}
	if response.Code != "0" {
		return nil, log.Errorf("error code [%v] message [%s]", response.Code, response.Msg)
	}
	//log.Json("response", response)
	for _, v := range response.Data {
		if v.InstType == types.SPOT {
			if len(ids) != 0 {
				ok := ids[v.InstId]
				if !ok {
					continue //ignore it
				}
			}
			prices = append(prices, v)
		}
	}
	return prices, nil
}

func (m *OkexClient) SpotTokens(ccys ...string) (tokens []types.TokenBase, err error) {
	var params = map[string]interface{}{
		"instType": types.SPOT,
	}
	var res *rest.RESTAPIResult
	res, err = m.client.Get(context.Background(), types.API_V5_ACCOUNT_INSTRUMENTS, &params)
	if err != nil {
		return nil, log.Errorf(err.Error())
	}
	log.Debugf("%s", res.Body)
	response := &types.InstrumentsResponseV5{}
	err = json.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		return nil, log.Errorf("response body json unmarshal error [%s]", err.Error())
	}
	if response.Code != "0" {
		return nil, log.Errorf("error code [%v] message [%s]", response.Code, response.Msg)
	}
	if len(ccys) != 0 {
		for _, ccy := range ccys {
			for _, tok := range response.Data {
				if tok.BaseCcy == ccy {
					tokens = append(tokens, tok)
				}
			}
		}
	} else {
		tokens = response.Data
	}
	return tokens, nil
}

func (m *OkexClient) SpotLoanTokens() (tokens []types.LoanToken, err error) {
	var params = map[string]interface{}{}
	var res *rest.RESTAPIResult
	res, err = m.client.Get(context.Background(), types.API_V5_ACCOUNT_INTEREST_RATE, &params)
	if err != nil {
		return nil, log.Errorf(err.Error())
	}
	log.Debugf("%s", res.Body)
	response := &types.LoanTokenResponseV5{}
	err = json.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		return nil, log.Errorf("response body json unmarshal error [%s]", err.Error())
	}
	if response.Code != "0" {
		return nil, log.Errorf("error code [%v] message [%s]", response.Code, response.Msg)
	}
	return response.Data, nil
}

func (m *OkexClient) SpotOrderBook(symbol string, level int) (depth types.OrderBook, err error) {
	var params = map[string]interface{}{
		"instId": symbol,
		"sz":     level,
	}
	var res *rest.RESTAPIResult
	res, err = m.client.Get(context.Background(), types.API_V5_MARKET_BOOKS, &params)
	if err != nil {
		return depth, log.Errorf(err.Error())
	}
	log.Debugf("%s", res.Body)
	response := &types.OrderBookResponseV5{}
	err = json.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		return depth, log.Errorf("response body json unmarshal error [%s]", err.Error())
	}
	if response.Code != "0" {
		return depth, log.Errorf("error code [%v] message [%s]", response.Code, response.Msg)
	}
	for _, d := range response.Data {
		for _, v := range d.Bids {
			depth.Bids = append(depth.Bids, types.DepthUnit{
				Price:         v[0],
				Quantity:      v[1],
				OrderQuantity: v[3],
			})
		}
		for _, v := range d.Asks {
			depth.Asks = append(depth.Asks, types.DepthUnit{
				Price:         v[0],
				Quantity:      v[1],
				OrderQuantity: v[3],
			})
		}
	}
	return depth, nil
}
