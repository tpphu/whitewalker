package cmd

import (
	"context"

	"../seed"

	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
	"go.uber.org/fx"
)

func handleSeedDB(lc fx.Lifecycle, db *gorm.DB) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			seed.NoteSeed(db)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			db.Close()
			return nil
		},
	})
}

func seedAction(appContext *cli.Context) {
	app := fx.New(
		fx.Provide(
			func() *cli.Context {
				return appContext
			},
			newLogger,
			newDB,
		),
		fx.Invoke(handleSeedDB),
	)

	app.Run()
}

// Seed is a definition of cli.Command used to migrate schema to database
var Seed = cli.Command{
	Name:  "seed",
	Usage: "Seed db",
	Action: func(appContext *cli.Context) error {
		seedAction(appContext)
		return nil
	},
}
