package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tpphu/whitewalker/repo"
	"github.com/urfave/cli"
)

// BuildEngine recieves cli.Context and return a gin.Engine
func BuildEngine(appContext *cli.Context) *gin.Engine {
	gin.SetMode(appContext.GlobalString("ginmode"))
	r := gin.Default()
	initDev(r)
	initNote(r)
	return r
}

func initNote(r *gin.Engine) {
	group := r.Group("/note")
	group.GET("/:id", func(c *gin.Context) {
		r := repo.NoteRepoImpl{}
		noteGetHanlder(c, r)
	})
}

func initDev(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
