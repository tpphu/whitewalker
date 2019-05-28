package handler

import (
	"github.com/gin-gonic/gin"
)

// BuildEngine recieves cli.Context and return a gin.Engine
func BuildEngine() *gin.Engine {
	r := gin.Default()
	InitDev(r)
	InitNote(r)
	return r
}

func InitNote(r *gin.Engine) {	
	group := r.Group("/note")
	group.GET("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func InitDev(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}