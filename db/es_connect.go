package db

import (
	"log"
	"context"

	"github.com/olivere/elastic/v7"

	"gin-exercise/yaml_config"
)


var EsClient *elastic.Client
var host = yaml_config.Cfg.Info.EsHosts

func init()  {
	var err error
	EsClient, err = elastic.NewClient(elastic.SetURL(yaml_config.Cfg.Info.EsHosts))
	if err != nil {
		log.Println(err)
		return
	}

	info, code, err := EsClient.Ping(yaml_config.Cfg.Info.EsHosts).Do(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("code:%d , %s\n", code, info.Version)
}
