package cmd

import (
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"github.com/urfave/cli"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
// https://github.com/uber-go/fx/blob/master/example_test.go

// newLogger create new logger
func newLogger() *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	return logger
}

func newDB(appContext *cli.Context) *gorm.DB {
	db, err := gorm.Open("mysql", appContext.GlobalString("database"))
	if err != nil {
		panic(err)
	}
	return db
}