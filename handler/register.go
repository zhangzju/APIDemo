package handler

import "github.com/gin-gonic/gin"

func Run(addr string) {
	r := gin.Default()
	r.GET("/ping", DecribeULB)
	r.Run(addr) // listen and serve on 0.0.0.0:8080
}
