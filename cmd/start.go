package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/tpphu/whitewalker/handler"
	"github.com/tpphu/whitewalker/server"
	"github.com/urfave/cli"
	"go.uber.org/fx"
)

// handleHTTPServer handles http server
func handleHTTPServer(lc fx.Lifecycle, appContext *cli.Context, logger *log.Logger, db *gorm.DB) {
	s := server.Server{
		Engine:  handler.BuildEngine(appContext, logger, db),
		Address: appContext.String("address"),
		Port:    appContext.String("port"),
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("!Start")
			// https://github.com/uber-go/fx/issues/627
			go s.Start(appContext)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("!Stop")
			s.Stop()
			return nil
		},
	})
}

// startAction starts command and init DI
func startAction(appContext *cli.Context) {
	app := fx.New(
		fx.Provide(
			func() *cli.Context {
				return appContext
			},
			newLogger,
			newDB,
		),
		fx.Invoke(handleHTTPServer),
	)

	app.Run()
}

// Start is a definition of cli.Command used to start http server
var Start = cli.Command{
	Name:  "start",
	Usage: "Start aplication",
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
		startAction(appContext)
		return nil
	},
}
