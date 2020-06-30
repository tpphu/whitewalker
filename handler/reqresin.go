package handler

import (
	"log"

	"github.com/imroc/req"
	"github.com/kataras/iris/v12"
	"github.com/tpphu/whitewalker/service"
)

type reqResInHandlerImpl struct {
	log           *log.Logger
	reqResService service.Service
}

func (handler reqResInHandlerImpl) inject(app *iris.Application) {
	group := app.Party("/reqresin/api/users")
	group.Get("/", handler.get)
}

func (handler reqResInHandlerImpl) get(c iris.Context) {
	param := req.Param{
		"p": 2,
	}
	res, _ := handler.reqResService.Get("/api/users", param)
	c.WriteString(res.String())
}
