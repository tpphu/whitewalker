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
	// Note handler
	noteHanler := &noteHandlerImpl{
		noteRepo: repo.NoteRepoImpl{
			DB: db,
		},
		log: logger,
	}
	initNote(r, noteHanler)
	// User handler
	userHanler := &userHandlerImpl{
		userRepo: repo.UserRepoImpl{
			DB: db,
		},
		log: logger,
	}
	initUser(r, userHanler)
	return r
}

func initNote(r *iris.Application, noteHanler *noteHandlerImpl) {
	group := r.Party("/note")
	group.Get("/{id:uint}", func(c iris.Context) {
		id := c.Params().GetUintDefault("id", 0)
		result, err := noteHanler.get(id)
		simpleReturnHandler(c, result, err)
	})
}

func initUser(r *iris.Application, userHanler *userHandlerImpl) {
	group := r.Party("/user")
	group.Get("/{id:uint}", func(c iris.Context) {
		id := c.Params().GetUintDefault("id", 0)
		result, err := userHanler.get(id)
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

func simpleReturnHandler(c iris.Context, result interface{}, err Error) {
	if err != nil {
		c.StatusCode(err.Status())
		c.JSON(iris.Map{
			"error": err.Error(),
		})
		return
	}
	c.JSON(result)
}
