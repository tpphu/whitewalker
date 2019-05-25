package main

import (
	"log"
	"os"

	"./cmd"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

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
