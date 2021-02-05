package main

import (
	"ginReal/common"
	"ginReal/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)


func main() {
	common.InitDB()

	r := gin.Default()
	routers.CollectRouter(r)
	port := viper.GetString("server.port")
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}

}
