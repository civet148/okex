package ws

import (
	"errors"
	"github.com/civet148/okex/types"
)

/*
	产品频道
*/
func (a *WsClient) PubInstruemnts(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(types.EVENT_BOOK_INSTRUMENTS, op, params, types.PERIOD_NONE, timeOut...)
}

func (a *WsClient) PubStatus(op string, timeOut ...int) (res bool, msg []*Msg, err error) {
	return a.PubChannel(types.EVENT_STATUS, op, nil, types.PERIOD_NONE, timeOut...)
}

/*
	行情频道
*/
func (a *WsClient) PubTickers(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(types.EVENT_BOOK_TICKERS, op, params, types.PERIOD_NONE, timeOut...)
}

/*
	持仓总量频道
*/
func (a *WsClient) PubOpenInsterest(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {
	return a.PubChannel(types.EVENT_BOOK_OPEN_INTEREST, op, params, types.PERIOD_NONE, timeOut...)
}

/*
	K线频道
*/
func (a *WsClient) PubKLine(op string, period types.Period, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(types.EVENT_BOOK_KLINE, op, params, period, timeOut...)
}

/*
	交易频道
*/
func (a *WsClient) PubTrade(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(types.EVENT_BOOK_TRADE, op, params, types.PERIOD_NONE, timeOut...)
}

/*
	预估交割/行权价格频道
*/
func (a *WsClient) PubEstDePrice(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(types.EVENT_BOOK_ESTIMATE_PRICE, op, params, types.PERIOD_NONE, timeOut...)

}

/*
	标记价格频道
*/
func (a *WsClient) PubMarkPrice(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(types.EVENT_BOOK_MARK_PRICE, op, params, types.PERIOD_NONE, timeOut...)
}

/*
	标记价格K线频道
*/
func (a *WsClient) PubMarkPriceCandle(op string, pd types.Period, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(types.EVENT_BOOK_MARK_PRICE_CANDLE_CHART, op, params, pd, timeOut...)
}

/*
	限价频道
*/
func (a *WsClient) PubLimitPrice(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(types.EVENT_BOOK_LIMIT_PRICE, op, params, types.PERIOD_NONE, timeOut...)
}

/*
	深度频道
*/
func (a *WsClient) PubOrderBooks(op string, channel string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	switch channel {
	// 400档快照
	case "books":
		return a.PubChannel(types.EVENT_BOOK_ORDER_BOOK, op, params, types.PERIOD_NONE, timeOut...)
	// 5档快照
	case "books5":
		return a.PubChannel(types.EVENT_BOOK_ORDER_BOOK5, op, params, types.PERIOD_NONE, timeOut...)
	// 400 tbt
	case "books-l2-tbt":
		return a.PubChannel(types.EVENT_BOOK_ORDER_BOOK_TBT, op, params, types.PERIOD_NONE, timeOut...)
	// 50 tbt
	case "books50-l2-tbt":
		return a.PubChannel(types.EVENT_BOOK_ORDER_BOOK50_TBT, op, params, types.PERIOD_NONE, timeOut...)

	default:
		err = errors.New("未知的channel")
		return
	}

}

/*
	期权定价频道
*/
func (a *WsClient) PubOptionSummary(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(types.EVENT_BOOK_OPTION_SUMMARY, op, params, types.PERIOD_NONE, timeOut...)
}

/*
	资金费率频道
*/
func (a *WsClient) PubFundRate(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(types.EVENT_BOOK_FUND_RATE, op, params, types.PERIOD_NONE, timeOut...)
}

/*
	指数K线频道
*/
func (a *WsClient) PubKLineIndex(op string, pd types.Period, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(types.EVENT_BOOK_KLINE_INDEX, op, params, pd, timeOut...)
}

/*
	指数行情频道
*/
func (a *WsClient) PubIndexTickers(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(types.EVENT_BOOK_INDEX_TICKERS, op, params, types.PERIOD_NONE, timeOut...)
}
