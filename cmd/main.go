package main

import (
	"fmt"
	"github.com/civet148/log"
	"github.com/civet148/okex"
	"github.com/civet148/okex/types"
	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
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
)

const (
	CMD_FLAG_NAME_API_ADDR    = "addr"
	CMD_FLAG_NAME_API_KEY     = "ak"
	CMD_FLAG_NAME_SECRET_KEY  = "sk"
	CMD_FLAG_NAME_PASS_PHRASE = "pass"
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
		log.Printf("\n\n")
		var tab *table.Table
		tab, err = gotable.Create("Token", "Total", "Available", "Frozen", "USD")
		for _, a := range balance.Data {
			for _, v := range a.Details {
				tab.AddRow([]string{v.Ccy, v.CashBal.Round(4).String(), v.AvailBal.Round(4).String(), v.FrozenBal.Round(4).String(), v.EqUsd.Round(4).String()})
			}
		}
		log.Printf(tab.String())
		return nil
	},
}
