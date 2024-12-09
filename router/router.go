package router

import (
	"github.com/gin-gonic/gin"
	"palamecia/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")
	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
	}
	return r
}
