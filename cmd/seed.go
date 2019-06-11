package cmd

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/tpphu/whitewalker/seed"
	"github.com/urfave/cli"
)

func seedAction(appContext *cli.Context) {
	db := newDB(appContext)
	defer db.Close()
	input := appContext.String("list")
	if input == "" {
		typeOfSeeder := reflect.TypeOf(seed.Seeder{})
		lenOfMethods := typeOfSeeder.NumMethod()
		fmt.Println("[1] You should pass argument --list * or -l *")
		fmt.Println("[2] Or you can pass one or more (use \",\" to separate) functions listed below:")
		for i := 0; i < lenOfMethods; i++ {
			fmt.Println("- " + typeOfSeeder.Method(i).Name)
		}
		return
	}
	seeder := seed.Seeder{
		DB: db,
	}
	s := reflect.ValueOf(seeder)
	arguments := []reflect.Value{}
	if input == "*" {
		lenOfMethods := s.NumMethod()
		for i := 0; i < lenOfMethods; i++ {
			s.Method(i).Call(arguments)
		}
		return
	}
	list := strings.Split(input, ",")
	for _, n := range list {
		s.MethodByName(n).Call(arguments)
	}
}

// Seed is a definition of cli.Command used to migrate schema to database
var Seed = cli.Command{
	Name:  "seed",
	Usage: "Seed db",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "list, l",
			Usage: "List the name of seed func should be run",
		},
	},
	Action: func(appContext *cli.Context) error {
		seedAction(appContext)
		return nil
	},
}
