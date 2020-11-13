package login

import (
	"errors"

	"gin-exercise/db"
	"gin-exercise/models"
)

func Login(username, password string) error {
	usertable := db.Dbtable("user")
	password = Addsalt(password)
	var user []*models.Userlogin
	row_num := usertable.Where(&models.Userlogin{Name: username, Password: password}).Find(&user).RowsAffected
	if row_num <= 0 {
		return errors.New("账号密码错误")
	}
	return nil

}
