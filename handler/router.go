package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tpphu/whitewalker/repo"
	"github.com/urfave/cli"
)

// BuildEngine recieves cli.Context and return a gin.Engine
func BuildEngine(appContext *cli.Context, logger *log.Logger, db *gorm.DB) *gin.Engine {
	gin.SetMode(appContext.GlobalString("ginmode"))
	r := gin.Default()
	initDev(r)
	initNote(r, logger, db)
	return r
}

func initNote(r *gin.Engine, logger *log.Logger, db *gorm.DB) {
	group := r.Group("/note")
	group.GET("/:id", func(c *gin.Context) {
		r := repo.NoteRepoImpl{
			DB: db,
		}
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
