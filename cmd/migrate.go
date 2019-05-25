package cmd

import (
	"github.com/urfave/cli"
)

// Migrate is a definition of cli.Command used to migrate schema to database
var Migrate cli.Command = cli.Command{
	Name:  "migrate",
	Usage: "Migrate db",
}
