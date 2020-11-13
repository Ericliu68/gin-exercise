package login

import (
	"fmt"
	"log"
	"time"

	"gin-exercise/db"
	"gin-exercise/models"
)

func Register(username, password string, age int) string {
	var users models.Userlogin
	db := db.Dbtable("user")
	row_num := db.Where("name = ?", username).Find(&users).RowsAffected
	if row_num > 0 {
		return fmt.Sprintf("%s", "账户已经存在了")
	}
	password = Addsalt(password)
	var user models.Userlogin
	user.Password = password
	user.Name = username
	user.Age = age
	time_now := time.Now()
	user.Create_time = &time_now

	tx := db.Begin()
	result := tx.Create(&user)
	if result.Error != nil {
		tx.Rollback()
		log.Printf("插入账户失败:: %v", result.Error)
		return "插入失败"
	} else {
		tx.Commit()
		return ""
	}

}

func CheckRole(username string) {

}
