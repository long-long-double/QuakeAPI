package utils

import (
	"QuakeAPI/log"
	"QuakeAPI/model"
	"gopkg.in/yaml.v2"
)

func ReadYaml(data string) model.YamlConfig {
	config := model.YamlConfig{}
	err := yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Log("Check Your Config Yaml", log.ERROR)
	}
	return config
}

func YamlToInput(config model.YamlConfig) model.Input {
	input := model.Input{}
	input.Total = config.Search.Total
	input.Search = config.Search.Query
	input.Output = config.Search.Output
	input.Email = config.Login.Email
	input.Key = config.Login.Key
	input.UserInfo = config.Login.Userinfo
	input.Quake = config.Use.Quake
	input.Fofa = config.Use.Fofa
	return input
}
