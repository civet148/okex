package main

import (
	"fmt"
	"github.com/civet148/log"
	"github.com/civet148/okex"
	"github.com/civet148/okex/types"
	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
	"github.com/shopspring/decimal"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"strings"
)

const (
	Version     = "v0.3.0"
	ProgramName = "okex"
)

var (
	BuildTime = "2024-05-13"
	GitCommit = ""
)

const (
	CMD_NAME_RUN       = "run"
	CMD_NAME_START     = "start"
	CMD_NAME_BALANCE   = "balance"
	CMD_NAME_BUY       = "buy"
	CMD_NAME_SELL      = "sell"
	CMD_NAME_LIST      = "list"
	CMD_NAME_CANCEL    = "cancel"
	CMD_NAME_PRICE     = "price"
	CMD_NAME_LIST_LOAN = "list-loan"
)

const (
	CMD_FLAG_NAME_DEBUG       = "debug"
	CMD_FLAG_NAME_CONFIG      = "config"
	CMD_FLAG_NAME_API_ADDR    = "addr"
	CMD_FLAG_NAME_API_KEY     = "ak"
	CMD_FLAG_NAME_SECRET_KEY  = "sk"
	CMD_FLAG_NAME_PASS_PHRASE = "pass"
	CMD_FLAG_NAME_ORDER_TYPE  = "order-type"
	CMD_FLAG_NAME_BASE        = "base"
	CMD_FLAG_NAME_PRICE       = "price"
	CMD_FLAG_NAME_ORDER_NO    = "order-no"
)

func init() {
	log.SetLevel("info")
}

func grace() {
	//capture signal of Ctrl+C and gracefully exit
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	go func() {
		for {
			select {
			case s := <-sigChannel:
				{
					if s != nil && s == os.Interrupt {
						fmt.Printf("Ctrl+C signal captured, program exiting...\n")
						close(sigChannel)
						os.Exit(0)
					}
				}
			}
		}
	}()
}

func main() {

	grace()

	local := []*cli.Command{
		runCmd,
		buyCmd,
		sellCmd,
		listCmd,
		cancelCmd,
		balanceCmd,
		priceCmd,
		listLoanCmd,
	}
	app := &cli.App{
		Name:        ProgramName,
		Version:     fmt.Sprintf("%s %s commit %s", Version, BuildTime, GitCommit),
		Description: "RESTFul & Websocket API V5 for OKEX",
		Flags:       []cli.Flag{},
		Commands:    local,
		Action:      nil,
	}
	if err := app.Run(os.Args); err != nil {
		log.Errorf("exit in error %s", err)
		os.Exit(1)
		return
	}
}

var apiFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    CMD_FLAG_NAME_DEBUG,
		Usage:   "debug request & response",
		Aliases: []string{"d"},
		EnvVars: []string{"OK_API_DEBUG"},
	},
	&cli.StringFlag{
		Name:    CMD_FLAG_NAME_CONFIG,
		Usage:   "config path of policy",
		Aliases: []string{"c"},
		Value:   types.OKEX_POLICY_CONFIG,
	},
	&cli.StringFlag{
		Name:    CMD_FLAG_NAME_API_ADDR,
		Usage:   "API server address of OKEX",
		Value:   types.OKEX_API_ADDR,
		Aliases: []string{"u"},
		EnvVars: []string{"OK_API_ADDR"},
	},
	&cli.StringFlag{
		Name:    CMD_FLAG_NAME_API_KEY,
		Usage:   "API key",
		Aliases: []string{"a"},
		EnvVars: []string{"OK_ACCESS_KEY"},
	},
	&cli.StringFlag{
		Name:    CMD_FLAG_NAME_SECRET_KEY,
		Usage:   "Secret key",
		Aliases: []string{"s"},
		EnvVars: []string{"OK_ACCESS_SECRET"},
	},
	&cli.StringFlag{
		Name:    CMD_FLAG_NAME_PASS_PHRASE,
		Usage:   "passphrase",
		Aliases: []string{"p"},
		EnvVars: []string{"OK_ACCESS_PASSPHRASE"},
	},
}

var runCmd = &cli.Command{
	Name:    CMD_NAME_RUN,
	Usage:   "run as a strategy service",
	Aliases: []string{CMD_NAME_START},
	Flags:   apiFlags,
	Action: func(cctx *cli.Context) error {
		return nil
	},
}

var balanceCmd = &cli.Command{
	Name:  CMD_NAME_BALANCE,
	Usage: "get account balance",
	Flags: apiFlags,
	Action: func(cctx *cli.Context) error {
		if cctx.IsSet(CMD_FLAG_NAME_DEBUG) {
			log.SetLevel("debug")
		}
		var strApiAddr = cctx.String(CMD_FLAG_NAME_API_ADDR)
		apiKeyInfo := &types.APIKeyInfo{
			ApiKey:     cctx.String(CMD_FLAG_NAME_API_KEY),
			PassPhrase: cctx.String(CMD_FLAG_NAME_PASS_PHRASE),
			SecKey:     cctx.String(CMD_FLAG_NAME_SECRET_KEY),
		}
		opts := &okex.Options{
			ApiUrl:         strApiAddr,
			TimeoutSeconds: 30,
			IsDebug:        cctx.Bool(CMD_FLAG_NAME_DEBUG),
			IsSimulate:     false,
		}
		client := okex.NewOkexClient(apiKeyInfo, opts)
		balance, err := client.Balance()
		if err != nil {
			return err
		}
		if cctx.Bool(CMD_FLAG_NAME_DEBUG) {
			log.Json(balance)
		}
		var tab *table.Table
		for _, a := range balance.Data {
			var strUSD = fmt.Sprintf("USD [%s]", a.TotalEq.Round(4).String())
			tab, err = gotable.Create("Token", "Total", "Available", "Frozen", strUSD, "Price")
			for _, v := range a.Details {
				if !v.CashBal.IsPositive() {
					continue
				}
				price := v.EqUsd.Div(v.CashBal).Round(8)
				tab.AddRow([]string{v.Ccy, v.CashBal.Round(4).String(), v.AvailBal.Round(4).String(), v.FrozenBal.Round(4).String(), v.EqUsd.Round(4).String(), price.String()})
			}
		}
		log.Printf(tab.String())
		return nil
	},
}

var tradeFlags = append(apiFlags, []cli.Flag{
	&cli.StringFlag{
		Name:  CMD_FLAG_NAME_ORDER_TYPE,
		Usage: "order type [limit/market/fok/ioc...]",
		Value: string(types.OrderType_Market),
	},
	&cli.StringFlag{
		Name:  CMD_FLAG_NAME_BASE,
		Usage: "base token",
		Value: types.TradeBaseUSDT,
	},
	&cli.Float64Flag{
		Name:  CMD_FLAG_NAME_PRICE,
		Usage: "trade price",
	},
	&cli.StringFlag{
		Name:  CMD_FLAG_NAME_ORDER_NO,
		Usage: "trade order no",
	},
}...)

var buyCmd = &cli.Command{
	Name:      CMD_NAME_BUY,
	Usage:     "spot buy",
	ArgsUsage: "<currency> <quantity>",
	Flags:     tradeFlags,
	Action: func(cctx *cli.Context) error {
		strOrderId, err := tradeOrder(cctx, types.TradeSide_Buy)
		if err != nil {
			return log.Errorf(err.Error())
		}
		log.Infof("order id [%s]", strOrderId)
		return nil
	},
}

var sellCmd = &cli.Command{
	Name:      CMD_NAME_SELL,
	Usage:     "spot sell",
	ArgsUsage: "<currency> <quantity>",
	Flags:     tradeFlags,
	Action: func(cctx *cli.Context) error {

		if cctx.IsSet(CMD_FLAG_NAME_DEBUG) {
			log.SetLevel("debug")
		}
		strOrderId, err := tradeOrder(cctx, types.TradeSide_Sell)
		if err != nil {
			return log.Errorf(err.Error())
		}
		log.Infof("order id [%s]", strOrderId)
		return nil
	},
}

func tradeOrder(cctx *cli.Context, side types.TradeSide) (orderId string, err error) {
	var strApiAddr = cctx.String(CMD_FLAG_NAME_API_ADDR)
	apiKeyInfo := &types.APIKeyInfo{
		ApiKey:     cctx.String(CMD_FLAG_NAME_API_KEY),
		PassPhrase: cctx.String(CMD_FLAG_NAME_PASS_PHRASE),
		SecKey:     cctx.String(CMD_FLAG_NAME_SECRET_KEY),
	}
	opts := &okex.Options{
		ApiUrl:         strApiAddr,
		TimeoutSeconds: 30,
		IsDebug:        cctx.Bool(CMD_FLAG_NAME_DEBUG),
		IsSimulate:     false,
	}
	client := okex.NewOkexClient(apiKeyInfo, opts)
	if cctx.Args().Len() != 2 {
		return "", fmt.Errorf("args required 2 but got %v", cctx.Args().Len())
	}
	ccy := strings.ToUpper(cctx.Args().Get(0))
	qty := cctx.Args().Get(1)
	quantity, err := decimal.NewFromString(qty)
	if err != nil {
		return "", fmt.Errorf("wrong quantity number %s, error [%s]", qty, err.Error())
	}
	base := cctx.String(CMD_FLAG_NAME_BASE)
	price := cctx.Float64(CMD_FLAG_NAME_PRICE)
	orderNo := cctx.String(CMD_FLAG_NAME_ORDER_NO)
	orderType := types.OrderType(cctx.String(CMD_FLAG_NAME_ORDER_TYPE))

	if (orderType == types.OrderType_Limit ||
		orderType == types.OrderType_Fok ||
		orderType == types.OrderType_Ioc) &&
		price == 0 {
		return "", fmt.Errorf("limit order price required")
	}

	req := types.TradeRequest{
		Side:      side,
		OrderType: orderType,
		TradeMode: types.TradeMode_Cash,
		Ccy:       ccy,
		Base:      base,
		Price:     decimal.NewFromFloat(price),
		Quantity:  quantity,
		OrderNo:   orderNo,
	}
	_ = client
	log.Json("trade request", req)
	return client.SpotTradeOrder(&req)
}

var listCmd = &cli.Command{
	Name:      CMD_NAME_LIST,
	Usage:     "list pending orders",
	ArgsUsage: "[PEPE_USDT,BTC_USDT...]",
	Flags:     tradeFlags,
	Action: func(cctx *cli.Context) error {

		if cctx.IsSet(CMD_FLAG_NAME_DEBUG) {
			log.SetLevel("debug")
		}
		var strApiAddr = cctx.String(CMD_FLAG_NAME_API_ADDR)
		apiKeyInfo := &types.APIKeyInfo{
			ApiKey:     cctx.String(CMD_FLAG_NAME_API_KEY),
			PassPhrase: cctx.String(CMD_FLAG_NAME_PASS_PHRASE),
			SecKey:     cctx.String(CMD_FLAG_NAME_SECRET_KEY),
		}
		var instIds []string
		instIds = strings.Split(cctx.Args().First(), ",")
		opts := &okex.Options{
			ApiUrl:         strApiAddr,
			TimeoutSeconds: 30,
			IsDebug:        cctx.Bool(CMD_FLAG_NAME_DEBUG),
			IsSimulate:     false,
		}
		client := okex.NewOkexClient(apiKeyInfo, opts)

		var strOrderType = string(types.OrderType_Limit)
		if cctx.IsSet(CMD_FLAG_NAME_ORDER_TYPE) {
			strOrderType = cctx.String(CMD_FLAG_NAME_ORDER_TYPE)
		}
		orders, err := client.SpotPendingOrders(strOrderType, instIds...)
		if err != nil {
			return log.Errorf(err.Error())
		}

		var tab *table.Table
		tab, err = gotable.Create("Instance ID", "Order ID", "Price", "Quantity", "Side", "USD")
		for _, v := range orders {
			usd := v.Sz.Mul(v.Px)
			err = tab.AddRow([]string{v.InstId, v.OrdId, v.Px.Round(8).String(), v.Sz.Round(8).String(), v.Side, usd.Round(5).String()})
			if err != nil {
				return log.Errorf(err.Error())
			}
		}
		log.Printf(tab.String())
		return nil
	},
}

var cancelCmd = &cli.Command{
	Name:      CMD_NAME_CANCEL,
	Usage:     "cancel pending order",
	ArgsUsage: "<inst id> <order id>",
	Flags:     tradeFlags,
	Action: func(cctx *cli.Context) error {
		if cctx.IsSet(CMD_FLAG_NAME_DEBUG) {
			log.SetLevel("debug")
		}
		var strApiAddr = cctx.String(CMD_FLAG_NAME_API_ADDR)
		apiKeyInfo := &types.APIKeyInfo{
			ApiKey:     cctx.String(CMD_FLAG_NAME_API_KEY),
			PassPhrase: cctx.String(CMD_FLAG_NAME_PASS_PHRASE),
			SecKey:     cctx.String(CMD_FLAG_NAME_SECRET_KEY),
		}
		var strInstId = cctx.Args().First()
		if strInstId == "" {
			return fmt.Errorf("inst id requires")
		}
		var strOrderId = cctx.Args().Get(1)
		if strOrderId == "" {
			return fmt.Errorf("order id requires")
		}
		opts := &okex.Options{
			ApiUrl:         strApiAddr,
			TimeoutSeconds: 30,
			IsDebug:        cctx.Bool(CMD_FLAG_NAME_DEBUG),
			IsSimulate:     false,
		}
		client := okex.NewOkexClient(apiKeyInfo, opts)
		if err := client.SpotCancelOrder(strInstId, strOrderId); err != nil {
			return log.Errorf(err.Error())
		}
		log.Infof("cancel inst id [%s] order id [%s] success", strInstId, strOrderId)
		return nil
	},
}

var priceCmd = &cli.Command{
	Name:      CMD_NAME_PRICE,
	Usage:     "query co-currency price",
	ArgsUsage: "<inst id>",
	Flags:     tradeFlags,
	Action: func(cctx *cli.Context) error {
		if cctx.IsSet(CMD_FLAG_NAME_DEBUG) {
			log.SetLevel("debug")
		}
		var strApiAddr = cctx.String(CMD_FLAG_NAME_API_ADDR)
		apiKeyInfo := &types.APIKeyInfo{
			ApiKey:     cctx.String(CMD_FLAG_NAME_API_KEY),
			PassPhrase: cctx.String(CMD_FLAG_NAME_PASS_PHRASE),
			SecKey:     cctx.String(CMD_FLAG_NAME_SECRET_KEY),
		}
		var strInstId = cctx.Args().First()
		opts := &okex.Options{
			ApiUrl:         strApiAddr,
			TimeoutSeconds: 30,
			IsDebug:        cctx.Bool(CMD_FLAG_NAME_DEBUG),
			IsSimulate:     false,
		}
		client := okex.NewOkexClient(apiKeyInfo, opts)
		if strInstId == "" {
			prices, err := client.SpotPrices()
			if err != nil {
				return log.Errorf(err.Error())
			}
			log.Json("spot prices", prices)
		} else {
			price, err := client.SpotPrice(strInstId)
			if err != nil {
				return log.Errorf(err.Error())
			}
			log.Infof("query inst id [%s] price [%s] success", strInstId, price.Last)
		}
		return nil
	},
}

var listLoanCmd = &cli.Command{
	Name:      CMD_NAME_LIST_LOAN,
	Usage:     "list loan tokens",
	ArgsUsage: "",
	Flags:     tradeFlags,
	Action: func(cctx *cli.Context) error {
		if cctx.IsSet(CMD_FLAG_NAME_DEBUG) {
			log.SetLevel("debug")
		}
		var strApiAddr = cctx.String(CMD_FLAG_NAME_API_ADDR)
		apiKeyInfo := &types.APIKeyInfo{
			ApiKey:     cctx.String(CMD_FLAG_NAME_API_KEY),
			PassPhrase: cctx.String(CMD_FLAG_NAME_PASS_PHRASE),
			SecKey:     cctx.String(CMD_FLAG_NAME_SECRET_KEY),
		}
		opts := &okex.Options{
			ApiUrl:         strApiAddr,
			TimeoutSeconds: 30,
			IsDebug:        cctx.Bool(CMD_FLAG_NAME_DEBUG),
			IsSimulate:     false,
		}
		client := okex.NewOkexClient(apiKeyInfo, opts)
		tokens, err := client.SpotLoanTokens()
		if err != nil {
			return log.Errorf(err.Error())
		}
		log.Json("loan list", tokens)
		return nil
	},
}
