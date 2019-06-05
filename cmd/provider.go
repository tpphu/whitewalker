package cmd

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/urfave/cli"
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
	db.LogMode(appContext.GlobalBool("dbdebug"))
	return db
}

func newRabbitMQ() {

}

func newRedis() {

}
