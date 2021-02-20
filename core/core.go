package core

import (
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

func DoQuake(input model.Input, mysqlConf db.MySQLConfig) {
	quakeCore := QuakeCore{}
	if input.UserInfo == true {
		quakeCore.GetUserInfo(input.Key)
	}
	if len(input.Search) != 0 && strings.TrimSpace(input.Search) != "" {
		doGetQuakeInfo(input, quakeCore, mysqlConf)
	}
}

func DoFofa(input model.Input, mysqlConf db.MySQLConfig) {
	fofaCore := FofaCore{}
	if input.UserInfo == true {
		fofaCore.GetUserInfo(input.Email, input.Key)
	}
	if len(input.Search) != 0 && strings.TrimSpace(input.Search) != "" {
		fofaCore.GetUserInfo(input.Email, input.Key)
		doGetFofaInfo(input, fofaCore, mysqlConf)
	}
}

func doGetFofaInfo(input model.Input, fofaCore FofaCore, mysqlConf db.MySQLConfig) {
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

func doGetQuakeInfo(input model.Input, quakeCore QuakeCore, mysqlConf db.MySQLConfig) {
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
