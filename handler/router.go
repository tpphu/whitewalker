package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

// BuildEngine recieves cli.Context and return a gin.Engine
func BuildEngine(c *cli.Context) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

func InitNote(c *cli.Context) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
