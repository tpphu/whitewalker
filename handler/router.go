package handler

import (
	"../repo"
	"github.com/gin-gonic/gin"
)


// BuildEngine recieves cli.Context and return a gin.Engine
func BuildEngine() *gin.Engine {
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