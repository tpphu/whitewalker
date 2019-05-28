package cmd

import (
	"github.com/jinzhu/gorm"
  "context"
  "log"
  "os"
  "time"
  "os/signal"

  "../handler"
  "../server"
  "go.uber.org/fx"
  "github.com/urfave/cli"
)

// handleHTTPServer handle http server
func handleHTTPServer(lc fx.Lifecycle, appContext *cli.Context, logger *log.Logger, db *gorm.DB) {
  s := server.Server{
    Engine:  handler.BuildEngine(),
    Address: appContext.String("address"),
    Port:    appContext.String("port"),
  }
  lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
      err := s.Start()
			return err
		},
		OnStop: func(ctx context.Context) error {
			err := s.Stop()
			return err
		},
	})
}

// startAction start command and init DI
func startAction(appContext *cli.Context) {
  app := fx.New(		
    fx.Provide(
      func() *cli.Context{
        return appContext
      },
      newLogger,
      newDB,
    ),
    fx.Invoke(handleHTTPServer),
  )

  startCtx := context.Background()
  if err := app.Start(startCtx); err != nil {
    log.Fatal(err)
  }

  // @Todo: This is not working
  quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
  <-quit
  
  stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
  defer cancel()
  if err := app.Stop(stopCtx); err != nil {
    log.Fatal(err)
  }
}

// Start is a definition of cli.Command used to start gin server
var Start = cli.Command{
  Name: "start",
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
