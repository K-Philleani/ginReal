package main

import (
	"ginReal/common"
	"ginReal/middleware"
	"ginReal/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)


func main() {
	common.InitDB()

	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	routers.CollectRouter(r)
	port := viper.GetString("server.port")
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
