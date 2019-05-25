package cmd

import (
	"../server"
	"github.com/urfave/cli"
)

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
			Address: c.String("address"),
			Port:    c.String("port"),
		}
		s.Start()
		return nil
	},
}
