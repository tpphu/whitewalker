package handler

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tpphu/whitewalker/repo"
	"github.com/urfave/cli"
)

// BuildEngine returns a *gin.Engine
func BuildEngine(appContext *cli.Context, logger *log.Logger, db *gorm.DB) *gin.Engine {
	gin.SetMode(appContext.GlobalString("ginmode"))
	r := gin.Default()
	initDev(r)
	initNote(r, logger, db)
	return r
}

func initNote(r *gin.Engine, logger *log.Logger, db *gorm.DB) {
	noteHanler := noteHandlerImpl{
		noteRepo: repo.NoteRepoImpl{
			DB: db,
		},
	}
	group := r.Group("/note")
	group.GET("/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, _ := strconv.Atoi(idParam)
		result, err := noteHanler.get(id)
		simpleReturnHandler(c, result, err)
	})
}

func initDev(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func simpleReturnHandler(c *gin.Context, result interface{}, err error) {
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, result)
}
