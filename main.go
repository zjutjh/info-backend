package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	"github.com/spf13/viper"
	"info/controller"
	"info/handler"
	"info/model"
)

func main() {
	// parse cmd args
	var options model.Options
	if _, err := flags.Parse(&options); err != nil {
		panic(err)
	}
	// read config from file(yaml）
	if options.ConfigPath != "" {
		viper.SetConfigFile(options.ConfigPath)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/etc/info")
		viper.AddConfigPath(".")
	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Config file 'config.yml' not found")
		} else {
			panic(err)
		}
	}

	// initial database connection
	handler.InitDB()

	// modify routers
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/info", controller.GetInfo)
		v1.GET("/dorm", controller.GetDorm)
	}

	// start server
	if err := router.Run(); err != nil {
		panic(err)
	}
}
