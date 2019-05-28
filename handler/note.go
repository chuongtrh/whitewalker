package handler

import (
	"github.com/tpphu/whitewalker/repo"
	"github.com/gin-gonic/gin"
)


func noteGetHanlder(c *gin.Context, repo repo.NoteRepo) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
