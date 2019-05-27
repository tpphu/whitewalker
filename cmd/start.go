package cmd

import (
	"log"
	"net/http"
	"os"

	"../handler"
	"../server"
	"github.com/urfave/cli"
)

// https://github.com/uber-go/fx/blob/master/example_test.go

// NewLogger create new logger
func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger.")
	return logger
}

// NewHandler create new handler
func NewHandler(logger *log.Logger) (http.Handler, error) {
	logger.Print("Executing NewHandler.")
	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		logger.Print("Got a request.")
	}), nil
}

// Start is a definition of cli.Command used to start gin server
var Start cli.Command = cli.Command{
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
	Action: func(c *cli.Context) error {
		s := server.Server{
			Engine:  handler.BuildEngine(c),
			Address: c.String("address"),
			Port:    c.String("port"),
		}
		s.Start()
		return nil
	},
}
