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
)

const (
	Version     = "v0.1.0"
	ProgramName = "okex"
)

var (
	BuildTime = "2024-04-19"
	GitCommit = ""
)

const (
	CMD_NAME_RUN     = "run"
	CMD_NAME_START   = "start"
	CMD_NAME_BALANCE = "balance"
	CMD_NAME_BUY     = "buy"
	CMD_NAME_SELL    = "sell"
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
		balanceCmd,
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
		Usage:   "debug raw data of interface response",
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
		//var strApiAddr = cctx.String(CMD_FLAG_NAME_API_ADDR)
		//apiKeyInfo := &types.APIKeyInfo{
		//	ApiKey:     cctx.String(CMD_FLAG_NAME_API_KEY),
		//	PassPhrase: cctx.String(CMD_FLAG_NAME_PASS_PHRASE),
		//	SecKey:     cctx.String(CMD_FLAG_NAME_SECRET_KEY),
		//}
		//client := okex.NewOkexClient(apiKeyInfo, strApiAddr)
		//balance, err := client.Balance()
		//if err != nil {
		//	return err
		//}
		//if cctx.Bool(CMD_FLAG_NAME_DEBUG) {
		//	log.Json(balance)
		//}
		//log.Printf("\n\n")
		//var tab *table.Table
		//for _, a := range balance.Data {
		//	var strUSD = fmt.Sprintf("USD[%s]", a.TotalEq.Round(4).String())
		//	tab, err = gotable.Create("Token", "Total", "Available", "Frozen", strUSD)
		//	for _, v := range a.Details {
		//		tab.AddRow([]string{v.Ccy, v.CashBal.Round(4).String(), v.AvailBal.Round(4).String(), v.FrozenBal.Round(4).String(), v.EqUsd.Round(4).String()})
		//	}
		//}
		//log.Printf(tab.String())
		return nil
	},
}

var balanceCmd = &cli.Command{
	Name:  CMD_NAME_BALANCE,
	Usage: "get account balance",
	Flags: apiFlags,
	Action: func(cctx *cli.Context) error {
		var strApiAddr = cctx.String(CMD_FLAG_NAME_API_ADDR)
		apiKeyInfo := &types.APIKeyInfo{
			ApiKey:     cctx.String(CMD_FLAG_NAME_API_KEY),
			PassPhrase: cctx.String(CMD_FLAG_NAME_PASS_PHRASE),
			SecKey:     cctx.String(CMD_FLAG_NAME_SECRET_KEY),
		}
		client := okex.NewOkexClient(apiKeyInfo, strApiAddr)
		balance, err := client.Balance()
		if err != nil {
			return err
		}
		if cctx.Bool(CMD_FLAG_NAME_DEBUG) {
			log.Json(balance)
		}
		log.Printf("\n\n")
		var tab *table.Table
		for _, a := range balance.Data {
			var strUSD = fmt.Sprintf("USD[%s]", a.TotalEq.Round(4).String())
			tab, err = gotable.Create("Token", "Total", "Available", "Frozen", strUSD)
			for _, v := range a.Details {
				tab.AddRow([]string{v.Ccy, v.CashBal.Round(4).String(), v.AvailBal.Round(4).String(), v.FrozenBal.Round(4).String(), v.EqUsd.Round(4).String()})
			}
		}
		log.Printf(tab.String())
		return nil
	},
}

var tradeFlags = append(apiFlags, []cli.Flag{
	&cli.StringFlag{
		Name:  CMD_FLAG_NAME_ORDER_TYPE,
		Usage: "order type [limit/market...]",
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
	client := okex.NewOkexClient(apiKeyInfo, strApiAddr)
	if cctx.Args().Len() != 2 {
		return "", fmt.Errorf("args required 2 but got %v", cctx.Args().Len())
	}
	ccy := cctx.Args().Get(0)
	qty := cctx.Args().Get(1)
	quantity, err := decimal.NewFromString(qty)
	if err != nil {
		return "", fmt.Errorf("wrong quantity number %s, error [%s]", qty, err.Error())
	}
	base := cctx.String(CMD_FLAG_NAME_BASE)
	price := cctx.Float64(CMD_FLAG_NAME_PRICE)
	orderNo := cctx.String(CMD_FLAG_NAME_ORDER_NO)
	orderType := types.OrderType(cctx.String(CMD_FLAG_NAME_ORDER_TYPE))

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
