package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"palamecia/global"
)

var DB *gorm.DB

func InitDB() {
	databaseConfig := AppConfig.Database

	var (
		db  *gorm.DB
		err error
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.DB_name,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//fmt.Println(os.Getwd())
	if err != nil {
		log.Fatal("failed to connect to database %v", err)
	}
	global.DB = db
}
