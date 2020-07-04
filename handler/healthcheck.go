package handler

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
)

type healthCheckHandlerImpl struct {
	log *log.Logger
	db  *gorm.DB
}

func (h healthCheckHandlerImpl) inject(app *iris.Application) {
	app.Get("/ping", h.ping)
}

func (h healthCheckHandlerImpl) ping(c iris.Context) {
	type Result struct {
		Ping string
	}
	var result Result
	h.db.Raw("SELECT 1 as ping").Scan(&result)
	c.JSON(iris.Map{
		"sucess":   true,
		"database": result.Ping == "1",
		"message":  "pong 4",
	})
}
