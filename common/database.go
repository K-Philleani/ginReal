package common

import (
	"fmt"
	"ginReal/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

var DB *gorm.DB

func InitConfig() {
	workDir, _ := os.Getwd()
	// 设置读取的文件名称
	viper.SetConfigName("application")
	// 设置文件类型
	viper.SetConfigType("yml")
	// 设置读取文件的路径
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func InitDB() *gorm.DB{
	InitConfig()
	var err error
	username, _ := Base64Decode([]byte(viper.GetString("datasource.username")))
	password, _ := Base64Decode([]byte(viper.GetString("datasource.password")))
	host, _ := Base64Decode([]byte(viper.GetString("datasource.host")))
	port, _ := Base64Decode([]byte(viper.GetString("datasource.port")))
	database, _ := Base64Decode([]byte(viper.GetString("datasource.database")))
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		username,
		password,
		host,
		port,
		database,
		charset,
		)
	fmt.Println(args)
	DB, err = gorm.Open(mysql.Open(args), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("fail to connect database, err:" + err.Error())
	}
	DB.AutoMigrate(&model.User{})
	log.Println("数据库连接成功")
	return DB
}

func GetDB() *gorm.DB {
	return DB
}


