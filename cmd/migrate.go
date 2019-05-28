package cmd

import (
  "os"
  "os/signal"
  "log"
  "time"
  "context"

  "../model"

	"go.uber.org/fx"
  "github.com/urfave/cli"
  "github.com/jinzhu/gorm"
)

func handleMigrateDB(db *gorm.DB) {
  db.AutoMigrate(&model.Note{})
}

func migrateAction(appContext *cli.Context) {
	app := fx.New(		
	  fx.Provide(
      func() *cli.Context{
        return appContext
      },
      newLogger,
      newDB,
	  ),
	  fx.Invoke(handleMigrateDB),
	)
  
	startCtx := context.Background()
	if err := app.Start(startCtx); err != nil {
	  log.Fatal(err)
	}
  
	// @Todo: This is not working
	quit := make(chan os.Signal)
	  signal.Notify(quit, os.Interrupt)
	<-quit
	
	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
	  log.Fatal(err)
	}
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
