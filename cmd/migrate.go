package cmd

import (
	"github.com/tpphu/whitewalker/model"

	"github.com/urfave/cli"
)

func migrateAction(appContext *cli.Context) {
	db := newDB(appContext)
	defer db.Close()
	db.AutoMigrate(&model.Note{}, &model.User{}, &model.Department{})
}

// Migrate is a definition of cli.Command used to migrate schema to database
var Migrate = cli.Command{
	Name:  "migrate",
	Usage: "Migrate db",
	Action: func(appContext *cli.Context) error {
		migrateAction(appContext)
		return nil
	},
}
