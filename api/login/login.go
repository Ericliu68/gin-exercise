package login

import (
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"

	"gin-user/models"
	"gin-user/utils"
)

// api 登录
func Login(c *gin.Context) {
	var login models.UserLogin
	// 判断名字和密码不能为空
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "请正确填写账户或密码"})
		return
	}
	// 判断帐号密码长度不能小于8
	if utf8.RuneCountInString(login.UserName) < 5 || utf8.RuneCountInString(login.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "帐号密码长度太小"})
		return
	}
	ok, err := utils.CheckPassword(login)
	if err != nil {
		utils.Logger.Info(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "服务器处理错误"})
		return
	}

	if ok {
		dbUser, err := utils.GetDbUser(login)
		token, err := utils.GetToken(login.UserName, dbUser.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "服务器处理错误"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "登录成功", "token": token})
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "帐号密码不正确"})
		return
	}
}
