package cmd

import (
	"github.com/urfave/cli"
)

// startAction starts command and init DI
func testAction(appContext *cli.Context) {
	db := newDB(appContext)
	defer db.Close()
}

// Test is a definition of cli.Command used to start http server
var Test = cli.Command{
	Name:  "test",
	Usage: "Test aplication",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "port, p",
			Value:  "8080",
			Usage:  "Listen on a port",
			EnvVar: "PORT",
		},
		cli.StringFlag{
			Name:   "address, a",
			Value:  "0.0.0.0",
			Usage:  "Bind to an address",
			EnvVar: "ADDRESS",
		},
	},
	Action: func(appContext *cli.Context) error {
		testAction(appContext)
		return nil
	},
}
