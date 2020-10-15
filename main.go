package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"ceshi":"hello world",
		})
	})

	r.Run(":8088")   //指定端口，默认8088
}
