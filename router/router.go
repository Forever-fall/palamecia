package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")
	hello := api.Group("/hello")
	hello.GET("", func(ctx *gin.Context) {
		fmt.Println("hello")
		ctx.JSON(200, gin.H{ // 返回一个 JSON 响应
			"message": "hello",
		})
	})
	return r
}
