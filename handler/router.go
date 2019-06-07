package handler

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/tpphu/whitewalker/repo"
	"github.com/urfave/cli"
)

// BuildEngine returns a *iris.Application
func BuildEngine(appContext *cli.Context, logger *log.Logger, db *gorm.DB) *iris.Application {
	r := iris.Default()
	r.Logger().SetLevel(appContext.GlobalString("loglevel"))
	initDev(r)
	initNote(r, logger, db)
	return r
}

func initNote(r *iris.Application, logger *log.Logger, db *gorm.DB) {
	noteHanler := noteHandlerImpl{
		noteRepo: repo.NoteRepoImpl{
			DB: db,
		},
	}
	group := r.Party("/note")
	group.Get("/{id:uint}", func(c iris.Context) {
		id := c.Params().GetUintDefault("id", 0)
		result, err := noteHanler.get(id)
		simpleReturnHandler(c, result, err)
	})
}

func initDev(r *iris.Application) {
	r.Get("/ping", func(c iris.Context) {
		c.JSON(iris.Map{
			"message": "pong",
		})
	})
}

func simpleReturnHandler(c iris.Context, result interface{}, err error) {
	if err != nil {
		c.StatusCode(400)
		c.JSON(iris.Map{
			"error": err.Error(),
		})
		return
	}
	c.JSON(result)
}
