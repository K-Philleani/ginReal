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
	authGroup := r.Group("/auth", middleware.AuthMiddleware())
	{
		authGroup.GET("/info", controller.Info)
		authGroup.GET("/user/list", controller.GetUserList)
		authGroup.POST("/user/delete", controller.DeleteUserByPhone)
	}
}
