package handler

import (
	"log"

	"github.com/kataras/iris"
	"github.com/tpphu/whitewalker/repo"
)

type noteHandlerImpl struct {
	noteRepo repo.NoteRepo
	log      *log.Logger
}

func (handler noteHandlerImpl) inject(app *iris.Application) {
	group := app.Party("/note")
	group.Get("/{id:uint}", handler.get)
}

func (handler noteHandlerImpl) get(c iris.Context) {
	id := c.Params().GetUintDefault("id", 0)
	note, err := handler.noteRepo.Find(id)
	if err != nil {
		simpleReturnHandler(c, note, NewNotFoundErr(err))
		return
	}
	simpleReturnHandler(c, note, nil)
}
