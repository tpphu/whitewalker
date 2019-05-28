package handler

import (
	"../repo"
	"github.com/gin-gonic/gin"
)


func noteGetHanlder(c *gin.Context, repo repo.NoteRepo) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
