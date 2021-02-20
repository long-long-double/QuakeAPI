package main

import (
	"QuakeAPI/core"
	"QuakeAPI/db"
	"QuakeAPI/log"
	"QuakeAPI/utils"
)

func main() {
	var mysqlConf db.MySQLConfig
	input := utils.ParseInput()
	if input.Config == true {
		filename := "config.yaml"
		if utils.FileExist(filename) == true {
			data := utils.ReadYamlFile(filename)
			config := utils.ReadYaml(data)
			if config.MySQL.Use == true {
				mysqlConf = db.MySQLConfig{
					Username: config.MySQL.Username,
					Password: config.MySQL.Password,
					Network:  "tcp",
					Server:   config.MySQL.Server,
					Port:     config.MySQL.Port,
					Database: "mysql",
				}
			}
			input = utils.YamlToInput(config)
		} else {
			utils.CreateYamlFile(filename)
			log.Log("Generate Config Yaml File", log.INFO)
			log.Log("Please Edit It", log.INFO)
			return
		}
	}
	if input.Quake == true {
		utils.PrintLogo("quake")
		core.DoQuake(input, mysqlConf)
	} else if input.Fofa == true {
		utils.PrintLogo("fofa")
		core.DoFofa(input, mysqlConf)
	} else {
		log.Log("Fofa or Quake ? You Should Chose One", log.ERROR)
		return
	}
}
