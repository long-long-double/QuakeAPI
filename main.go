package main

import (
	"QuakeAPI/core"
	"QuakeAPI/log"
	"QuakeAPI/model"
	"QuakeAPI/utils"
	"bytes"
	"strings"
)

func main() {
	input := utils.ParseInput()
	if input.Quake == true {
		utils.PrintLogo("quake")
		doQuake(input)
	} else if input.Fofa == true {
		utils.PrintLogo("fofa")
		doFofa(input)
	} else {
		log.Log("Fofa or Quake ? You Should Chose One(--fofa/--quake)", log.ERROR)
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
	buffer := bytes.Buffer{}
	log.Log("Search : "+input.Search, log.INFO)
	if input.Total > 100 {
		var currentPage int
		var size int
		results, currentPage, size = fofaCore.GetSearchInfo(input.Search, input.Total, 1)
		buffer.WriteString(results)
		for i := 0; i < size/input.Total+1; i++ {
			results, currentPage, size = fofaCore.GetSearchInfo(input.Search, input.Total, currentPage+1)
			buffer.WriteString(results)
		}
	} else {
		results, _, _ = fofaCore.GetSearchInfo(input.Search, input.Total, 1)
		buffer.WriteString(results)
	}
	results = buffer.String()
	utils.WriteOutput(results, input.Output)
}

func doGetQuakeInfo(input model.Input, quakeCore core.QuakeCore) {
	var results string
	buffer := bytes.Buffer{}
	if input.Total > 100 {
		index := input.Total / 100
		pid, result := quakeCore.GetServiceInfo(input.Key, input.Search, 100, "")
		buffer.WriteString(result)
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
				buffer.WriteString(data)
			case <-quitChan:
				flag += 1
				if flag == index {
					goto finish
				}
			}
		}
	finish:
		results = buffer.String()
	} else {
		_, results = quakeCore.GetServiceInfo(input.Key, input.Search, input.Total, "")
	}
	utils.WriteOutput(results, input.Output)
}
