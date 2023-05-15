package database

import (
	"fmt"
	"log"
	"todolistapi/app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg config.DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Mysql_User, cfg.Mysql_Password, cfg.Mysql_Host, cfg.Mysql_Port, cfg.Mysql_DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Database connection error : ", err.Error())
		return nil
	}

	return db
}

func Migrate(db *gorm.DB) {
}