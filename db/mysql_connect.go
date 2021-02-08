package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gin-user/settings"
)

var DB *gorm.DB

// 初始化数据库连接
func DbInit() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		settings.Setting.Mysql.Username,
		settings.Setting.Mysql.Password,
		settings.Setting.Mysql.Path,
		settings.Setting.Mysql.DbName,
		settings.Setting.Mysql.Config,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		os.Exit(1)
		return

	}
	sql, err := DB.DB()
	sql.SetMaxIdleConns(settings.Setting.Mysql.MaxIdleConns)
	sql.SetMaxOpenConns(settings.Setting.Mysql.MaxOpenConns)
	sql.SetConnMaxLifetime(time.Hour)
}

func DbTable(table string) *gorm.DB {
	return DB.Table(table)
}
