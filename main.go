package main

import (
	"log"
	"os"

	"github.com/tpphu/whitewalker/cmd"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := cli.NewApp()

	app.Name = "whitewalker"
	app.Usage = "A Gin Application"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "database, db",
			Value:  "default:secret@/notes?charset=utf8&parseTime=True&loc=Local",
			Usage:  "Database connection",
			EnvVar: "DATABASE",
		},
		cli.BoolFlag{
			Name:   "dbdebug",
			Usage:  "Database debug",
			EnvVar: "DB_DEBUG",
		},
		cli.StringFlag{
			Name:   "ginmode",
			Value:  "release",
			Usage:  "Database debug",
			EnvVar: "GIN_MODE",
		},
	}

	app.Commands = []cli.Command{
		cmd.Start,
		cmd.Migrate,
		cmd.Seed,
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
