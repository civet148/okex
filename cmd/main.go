package main

import (
	"fmt"
	"github.com/civet148/log"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
)

const (
	Version     = "OKEX RESTFul API v5.x"
	ProgramName = "okex"
)

var (
	BuildTime = "2024-04-09"
	GitCommit = ""
)

const (
	CMD_NAME_BALANCE = "balance"
)

const (
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
		Name:     ProgramName,
		Version:  fmt.Sprintf("%s %s commit %s", Version, BuildTime, GitCommit),
		Flags:    []cli.Flag{},
		Commands: local,
		Action:   nil,
	}
	if err := app.Run(os.Args); err != nil {
		log.Errorf("exit in error %s", err)
		os.Exit(1)
		return
	}
}

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    CMD_FLAG_NAME_API_KEY,
		Usage:   "API key",
		Aliases: []string{"a"},
		EnvVars: []string{"OK-ACCESS-KEY"},
	},
	&cli.StringFlag{
		Name:    CMD_FLAG_NAME_SECRET_KEY,
		Usage:   "Secret key",
		Aliases: []string{"s"},
		EnvVars: []string{"OK-ACCESS-SECRET"},
	},
	&cli.StringFlag{
		Name:    CMD_FLAG_NAME_PASS_PHRASE,
		Usage:   "passphrase",
		Aliases: []string{"p"},
		EnvVars: []string{"OK-ACCESS-PASSPHRASE"},
	},
}

var balanceCmd = &cli.Command{
	Name:  CMD_NAME_BALANCE,
	Usage: "get account balance",
	Flags: flags,
	Action: func(cctx *cli.Context) error {

		return nil
	},
}
