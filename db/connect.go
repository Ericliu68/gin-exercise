package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := "root:9090@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("connect db err :: ", err)
		os.Exit(1)
	}

	sql, err := DB.DB()
	sql.SetConnMaxLifetime(time.Hour)
	sql.SetMaxOpenConns(100)
	sql.SetMaxIdleConns(10)

}

func Dbtable(table string) *gorm.DB {
	return DB.Table(table)
}
