package jwt_token

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type MyClaims struct {
	Username string
	jwt.StandardClaims
}

var TimeExpiration = time.Hour * 24

// get token
func GetToken(username string) (string, error) {
	c := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TimeExpiration).Unix(), // 过期时间
			Issuer:    "eric.liu",                            // 签发人
		},
	}
	privateKey, err := ioutil.ReadFile("./pem/private_key.pem")
	if err != nil {
		return "", fmt.Errorf("error reading private key file: %v\n", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", fmt.Errorf("error parsing RSA private key: %v\n", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	return token.SignedString(key)

}

// 解析JWT
func ParseToken(tokenString string) (bool, string) {
	publicKey, err := ioutil.ReadFile("./pem/public_key.pem")
	if err != nil {
		log.Printf("读取public_key出错::%v", err)
		return false, fmt.Sprintf("读取public_key出错::%v\n", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		log.Printf("error parsing RSA public key::%v\n", err)
		return false, fmt.Sprintf("error parsing RSA public key::%v\n", err)
	}

	parsedToken, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return key, nil
	})
	if err != nil {
		log.Printf("error parsing token:: %v\n", err)
		return false, fmt.Sprintf("error parsing token:: %v\n", err)
	}
	if claims, ok := parsedToken.Claims.(*MyClaims); ok && parsedToken.Valid {
		log.Println(claims)
		return true, ""
	}
	return false, ""
}

// get jwt-token from header
func GetJwtFromHeader(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请求头中auth为空",
		})
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请求头中auth格式有误",
		})
		c.Abort()
		return
	}
	mc, err := ParseToken(parts[1])
	if !mc {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "无效的Token",
		})
		c.Abort()
		return
	}
	// 将当前请求的username信息保存到请求的上下文c上
	//c.Set("username", mc.Username)
	c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
}
