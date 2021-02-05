package routers

import (
	"ginReal/controller"
	"ginReal/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/register", controller.Register)
		apiGroup.POST("/login", controller.Login)
	}
	authGroup := r.Group("/auth")
	{
		authGroup.GET("/info", middleware.AuthMiddleware() ,controller.Info)
	}
}
