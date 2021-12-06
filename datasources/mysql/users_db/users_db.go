package users_db

import (
	"fmt"
	"go-apk-users/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	Client *gorm.DB
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MysqlUsername, config.MysqlPassword, config.MySqlHost, config.MySqlTable,
	)
	var err error
	Client, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("database connected successfully")
}
