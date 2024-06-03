package types

const (
	API_V5_ACCOUNT_BALANCE       = "/api/v5/account/balance"       //余额查询
	API_V5_ACCOUNT_INSTRUMENTS   = "/api/v5/account/instruments"   //查询所有代币基础信息
	API_V5_TRADE_ORDER           = "/api/v5/trade/order"           //订单交易
	API_V5_PENDING_ORDERS        = "/api/v5/trade/orders-pending"  //查询挂单
	API_V5_CANCEL_ORDER          = "/api/v5/trade/cancel-order"    //取消挂单
	API_V5_MARKET_TICKER         = "/api/v5/market/ticker"         //单个产品行情信息
	API_V5_MARKET_TICKERS        = "/api/v5/market/tickers"        //全部产品行情信息
	API_V5_ACCOUNT_INTEREST_RATE = "/api/v5/account/interest-rate" //借币币种和利息
)

const (
	SPOT = "SPOT" //现货
)
