package cmd

import (
	"../model"

	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
	"go.uber.org/fx"
)

func handleMigrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.Note{})
}

func migrateAction(appContext *cli.Context) {
	app := fx.New(
		fx.Provide(
			func() *cli.Context {
				return appContext
			},
			newLogger,
			newDB,
		),
		fx.Invoke(handleMigrateDB),
	)

	app.Run()
}

// Migrate is a definition of cli.Command used to migrate schema to database
var Migrate cli.Command = cli.Command{
	Name:  "migrate",
	Usage: "Migrate db",
	Action: func(appContext *cli.Context) error {
		migrateAction(appContext)
		return nil
	},
}
