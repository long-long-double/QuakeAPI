package utils

import (
	"QuakeAPI/log"
	"io/ioutil"
	"os"
)

var template = `login:
  email: "your@email.com"       # 如果使用Fofa需要提供Email
  key: "your-key"               # 需要提供API Key
  userinfo: false               # 是否查询用户相关的信息

search:
  query: "service:http"         # 查询字符串
  output: "result.txt"          # 输出文件
  total: 1000                   # 查询个数（建议100的倍数）

use:
  quake: false                  # 是否使用quake引擎
  fofa: true                    # 是否使用fofa引擎

`

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	return false
}

func WriteOutput(content string, filename string) {
	if FileExist(filename) == true {
		log.Log("Output File Exist", log.INFO)
		err := os.Remove(filename)
		log.Log("Delete Old File", log.INFO)
		if err != nil {
			log.Log("Delete Output File Error:"+err.Error(), log.ERROR)
			return
		}
	}
	log.Log("Create Output File", log.INFO)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		log.Log("Create File Error:"+err.Error(), log.ERROR)
		return
	}
	_, err = file.Write([]byte(content))
	if err != nil {
		log.Log("Write File Error:"+err.Error(), log.ERROR)
		return
	}
}

func CreateYamlFile(filename string) {
	log.Log("Create Config File", log.INFO)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		log.Log("Create Config File Error:"+err.Error(), log.ERROR)
		return
	}
	_, err = file.Write([]byte(template))
	if err != nil {
		log.Log("Write Config File Error:"+err.Error(), log.ERROR)
		return
	}
}

func ReadYamlFile(filename string) string {
	bytesContent, err := ioutil.ReadFile(filename)
	if err == nil {
		return string(bytesContent)
	} else {
		log.Log("Read Config File Error", log.ERROR)
	}
	return ""
}
