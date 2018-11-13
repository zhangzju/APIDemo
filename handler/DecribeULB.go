package handler

import "github.com/gin-gonic/gin"

func DecribeULB(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
