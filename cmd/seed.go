package cmd

import (
	"github.com/urfave/cli"
)

// Seed is a definition of cli.Command used to seed data to database
var Seed = cli.Command{
	Name:  "seed",
	Usage: "Seed to db",
}
