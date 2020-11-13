package yaml_config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"gin-exercise/models"
)

var Cfg models.SettingYaml

func init() {
	ReadYamlConfig("/home/liu/workplace/gitlab/lion/services/user/config/config.development.yaml")
}
func ReadYamlConfig(path string) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &Cfg)
	if err != nil {
		log.Println(err.Error())
	}
}
