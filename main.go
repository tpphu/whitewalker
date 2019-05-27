package main

import (
	"log"
	"os"

	"./cmd"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "whitewalker"
	app.Usage = "A Gin Application"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		cmd.Start,
		cmd.Migrate,
		cmd.Seed,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
