package handler

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"info/data"
	"info/model"
)

func InitDB() {
	user := viper.GetString("username")
	passwd := viper.GetString("password")
	database := viper.GetString("database")
	hostname := viper.GetString("hostname")
	port := viper.GetString("port")
	if user == "" || passwd == "" || database == "" {
		panic("Invalid config")
	}
	var dsn string
	if hostname == "" || port == "" {
		dsn = fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, database)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, hostname, port, database)
	}
	var err error
	data.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := data.DB.AutoMigrate(&model.Student{}); err != nil {
		panic(err)
	}
}