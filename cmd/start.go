package cmd

import (
  "context"
  "log"
  "os"
  "time"

  "../handler"
  "../server"
  "go.uber.org/fx"
  "github.com/urfave/cli"
)

// https://github.com/uber-go/fx/blob/master/example_test.go

// newLogger create new logger
func newLogger() *log.Logger {
  logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
  return logger
}

// newLogger start server
func newInvoker(appContext *cli.Context, logger *log.Logger) {
  s := server.Server{
    Engine:  handler.BuildEngine(),
    Address: appContext.String("address"),
    Port:    appContext.String("port"),
  }
  s.Start()
}

// startAction start command and init DI
func startAction(appContext *cli.Context) {
  app := fx.New(		
    fx.Provide(
      func() *cli.Context{
        return appContext
      },
      newLogger,
    ),
    fx.Invoke(newInvoker),
  )

  // In a typical application, we could just use app.Run() here. Since we
  // don't want this example to run forever, we'll use the more-explicit Start
  // and Stop.
  startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
  defer cancel()
  if err := app.Start(startCtx); err != nil {
    log.Fatal(err)
  }

  // stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
  // defer cancel()
  // if err := app.Stop(stopCtx); err != nil {
  //   log.Fatal(err)
  // }
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
