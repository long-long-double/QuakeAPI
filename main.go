package main

import (
	"QuakeAPI/core"
	"QuakeAPI/db"
	"QuakeAPI/log"
	"QuakeAPI/model"
	"QuakeAPI/utils"
	"bytes"
	"strings"
	"sync"
)

type DataLock struct {
	Lock   sync.Mutex
	buffer bytes.Buffer
}

var QuitLock sync.Mutex
var lock = DataLock{sync.Mutex{}, bytes.Buffer{}}

var mysqlConf db.MySQLConfig

func main() {
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
		doQuake(input)
	} else if input.Fofa == true {
		utils.PrintLogo("fofa")
		doFofa(input)
	} else {
		log.Log("Fofa or Quake ? You Should Chose One", log.ERROR)
		return
	}
}

func doQuake(input model.Input) {
	quakeCore := core.QuakeCore{}
	if input.UserInfo == true {
		quakeCore.GetUserInfo(input.Key)
	}
	if len(input.Search) != 0 && strings.TrimSpace(input.Search) != "" {
		doGetQuakeInfo(input, quakeCore)
	}
}

func doFofa(input model.Input) {
	fofaCore := core.FofaCore{}
	if input.UserInfo == true {
		fofaCore.GetUserInfo(input.Email, input.Key)
	}
	if len(input.Search) != 0 && strings.TrimSpace(input.Search) != "" {
		fofaCore.GetUserInfo(input.Email, input.Key)
		doGetFofaInfo(input, fofaCore)
	}
}

func doGetFofaInfo(input model.Input, fofaCore core.FofaCore) {
	var results string
	log.Log("Search : "+input.Search, log.INFO)
	if input.Total > 100 {
		pageNum := input.Total/100 + 1
		dataChan := make(chan string)
		quitChan := make(chan bool, pageNum)
		lock.buffer.WriteString(results)
		for i := 1; i <= pageNum; i++ {
			currentPage := i
			go func() {
				results, _, _ = fofaCore.GetSearchInfo(input.Search, currentPage)
				dataChan <- results
				quitChan <- true
			}()
		}
		flag := 0
		for {
			select {
			case data := <-dataChan:
				lock.Lock.Lock()
				lock.buffer.WriteString(data)
				lock.Lock.Unlock()
			case <-quitChan:
				QuitLock.Lock()
				flag += 1
				QuitLock.Unlock()
				if flag == pageNum {
					goto finish
				}
			}
		}
	finish:
		results = lock.buffer.String()
	} else {
		results, _, _ = fofaCore.GetSearchInfo(input.Search, 1)
	}
	if input.Output == "save-to-mysql" {
		dataList := strings.Split(results, "\n")
		db.SaveDataToMySQL(dataList, mysqlConf)
	} else {
		utils.WriteOutput(results, input.Output)
	}
}

func doGetQuakeInfo(input model.Input, quakeCore core.QuakeCore) {
	var results string
	if input.Total > 100 {
		index := input.Total / 100
		pid, result := quakeCore.GetServiceInfo(input.Key, input.Search, 100, "")
		lock.buffer.WriteString(result)
		dataChan := make(chan string)
		quitChan := make(chan bool, index)
		for i := 0; i < index; i++ {
			go func() {
				pid, result = quakeCore.GetServiceInfo(input.Key, input.Search, 100, pid)
				dataChan <- result
				quitChan <- true
			}()
		}
		flag := 0
		for {
			select {
			case data := <-dataChan:
				lock.Lock.Lock()
				lock.buffer.WriteString(data)
				lock.Lock.Unlock()
			case <-quitChan:
				QuitLock.Lock()
				flag += 1
				QuitLock.Unlock()
				if flag == index {
					goto finish
				}
			}
		}
	finish:
		results = lock.buffer.String()
	} else {
		_, results = quakeCore.GetServiceInfo(input.Key, input.Search, input.Total, "")
	}
	if input.Output == "save-to-mysql" {
		dataList := strings.Split(results, "\n")
		db.SaveDataToMySQL(dataList, mysqlConf)
	} else {
		utils.WriteOutput(results, input.Output)
	}
}
