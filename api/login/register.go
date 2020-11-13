package login

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-exercise/login"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	if username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名不能为空",
		})
		return
	}
	password := c.PostForm("password")
	if password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "密码不能为空",
		})
		return
	}
	var age_num int
	var err error
	age := c.DefaultPostForm("age", "0")
	if age == "" {
		age_num = 0
	} else {
		age_num, err = strconv.Atoi(age)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "age解析失败",
			})
			return
		}
	}

	err_str := login.Register(username, password, age_num)
	if err_str != "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err_str,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "插入成功",
		})
	}

}
