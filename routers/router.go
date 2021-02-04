package routers

import (
	"ginReal/controller"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/register", controller.Register)
		apiGroup.POST("/login", controller.Login)
	}
}
