package cmd

import (
	"context"

	"github.com/tpphu/whitewalker/model"

	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
	"go.uber.org/fx"
)

func handleMigrateDB(lc fx.Lifecycle, db *gorm.DB) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			db.AutoMigrate(&model.Note{}, &model.User{}, &model.Department{})
			return nil
		},
		OnStop: func(ctx context.Context) error {
			db.Close()
			return nil
		},
	})
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
var Migrate = cli.Command{
	Name:  "migrate",
	Usage: "Migrate db",
	Action: func(appContext *cli.Context) error {
		migrateAction(appContext)
		return nil
	},
}
