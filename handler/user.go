package handler

import (
	"log"

	"github.com/kataras/iris"
	"github.com/tpphu/whitewalker/repo"
)

type userHandlerImpl struct {
	userRepo repo.UserRepo
	log      *log.Logger
}

func (handler userHandlerImpl) inject(app *iris.Application) {
	group := app.Party("/user")
	group.Get("/{id:uint}", handler.get)
}

func (handler userHandlerImpl) get(c iris.Context) {
	id := c.Params().GetUintDefault("id", 0)
	user, err := handler.userRepo.Find(id)
	if err != nil {
		simpleReturnHandler(c, user, NewNotFoundErr(err))
		return
	}
	simpleReturnHandler(c, user, nil)
}
