package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tpphu/whitewalker/repo"
)

func noteGetHanlder(c *gin.Context, repo repo.NoteRepo) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	note, err := repo.Find(id)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "not found",
		})
		return
	}
	c.JSON(200, note)
}
