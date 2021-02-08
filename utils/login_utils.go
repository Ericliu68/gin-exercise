package utils

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/scrypt"

	"gin-user/db"
	"gin-user/models"
	"gin-user/settings"
)

// 密码加盐
func GetEncryption(password string) (string, error) {
	dk, err := scrypt.Key([]byte(password), []byte(settings.Setting.Salt), 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(dk), nil
}

// 比对密码
func CheckPassword(user models.UserLogin) (bool, error) {
	password, err := GetEncryption(user.Password)
	if err != nil {
		return false, err
	}
	dbUser, err := GetDbUser(user)
	if err != nil {
		return false, err
	}

	if password == dbUser.Password {
		return true, nil
	} else {
		return false, nil
	}

}

// 数据库获取password
func GetDbUser(user models.UserLogin) (models.Users, error) {
	var dbUser models.Users
	//row_data := db.DbTable("users").Where("username = ?", user.UserName).First(&dbUser)
	row := db.DbTable("users").Where("name = ?", user.UserName).First(&dbUser).RowsAffected
	Logger.Info(user.UserName)
	if row == 1 {
		return dbUser, nil
	} else {
		return dbUser, fmt.Errorf("不存在此账户")
	}
}
