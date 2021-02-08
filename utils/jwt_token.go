package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type MyClaims struct {
	Username string
	Role     string
	jwt.StandardClaims
}

var TimeExpiration = time.Hour * 24

// get token
func GetToken(username, role string) (string, error) {
	c := MyClaims{
		username,
		role,
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
		return false, fmt.Sprintf("读取public_key出错::%v\n", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
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
			"success": false,
			"message": "请求头中auth为空",
		})
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "请求头中auth格式有误",
		})
		c.Abort()
		return
	}
	mc, err := ParseToken(parts[1])
	if !mc {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无效的Token",
		})
		c.Abort()
		return
	}
	// 将当前请求的username信息保存到请求的上下文c上
	//c.Set("username", mc.Username)
	c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
}

func PerssionToken(c *gin.Context) {

}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 创建路径
func CreatePath(path string) bool {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		Logger.Error(err.Error())
		return false
	}
	return true
}

func CreatePemKey() {

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("./pem/private_key.pem")
	if err != nil {
		panic(err)
	}
	err = pem.Encode(file, block)
	if err != nil {
		panic(err)
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("./pem/public_key.pem")
	if err != nil {
		panic(err)
	}
	err = pem.Encode(file, block)
	if err != nil {
		panic(err)
	}
	//panic(err)
}

func CreatePathPemKey() {
	Logger.Info("start")
	ok := Exists("./pem")
	if !ok {
		ok = CreatePath("./pem")
		if !ok {
			panic("初始化公私钥目录失败")
		}
		CreatePemKey()
	} else {
		ok1 := Exists("./pem/private_key.pem")
		ok2 := Exists("./pem/public_key.pem")
		if ok1 && ok2 {
			Logger.Info("ok")
		} else {
			CreatePemKey()
			Logger.Info("创建pem成功")
		}
	}

}
