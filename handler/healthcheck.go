package handler

import (
	"log"

	"github.com/kataras/iris"
)

type healthCheckHandlerImpl struct {
	log *log.Logger
}

func (handler healthCheckHandlerImpl) inject(app *iris.Application) {
	app.Get("/ping", handler.ping)
}

func (handler healthCheckHandlerImpl) ping(c iris.Context) {
	c.JSON(iris.Map{
		"message": "pong 11",
	})
}
