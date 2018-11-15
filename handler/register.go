package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "APIDemo/handler/docs"
)

// @title ULB Online API
// @version 0.0.1
// @description  This is a sample server Petstore server.
// @BasePath /api/v1/
func Run(addr string) {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", DecribeULB)
	r.Run(addr) // listen and serve on 0.0.0.0:8080
}
