package main

import (
	"github.com/gin-gonic/gin"

	"gin-exercise/api/login"
	"gin-exercise/jwt_token"
)

func main() {
	router := gin.Default()

	login_server := router.Group("/login_services/api")
	login_server.POST("/login", login.Login)
	login_server.POST("/register", jwt_token.GetJwtFromHeader, login.Register)

	router.Run("0.0.0.0:8080")
}

//import (
//	"log"
//	"context"
//
//	"github.com/olivere/elastic/v7"
//
//	"gin-exercise/yaml_config"
//)
//
//
//var EsClient *elastic.Client
//var host = yaml_config.Cfg.Info.EsHosts
//
//func init()  {
//	var err error
//	EsClient, err = elastic.NewClient(elastic.SetURL(yaml_config.Cfg.Info.EsHosts))
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	info, code, err := EsClient.Ping(yaml_config.Cfg.Info.EsHosts).Do(context.Background())
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//	log.Printf("code:%d , %s\n", code, info.Version)
//}
//
//func main(){
//
//}