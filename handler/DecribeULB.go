package handler

import "github.com/gin-gonic/gin"

// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"

func DecribeULB(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
