package login

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-exercise/jwt_token"
	"gin-exercise/login"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "缺少必要参数",
		})
		return
	}
	if password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "缺少必要参数",
		})
		return
	}

	err := login.Login(username, password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  fmt.Sprintf("%v", err),
		})
		return
	}

	token, err := jwt_token.GetToken(username)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功",
		"data": gin.H{
			"JWT":      token,
			"username": username,
		},
	})

}
