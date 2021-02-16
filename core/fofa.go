package core

import (
	"QuakeAPI/log"
	"QuakeAPI/model"
	"bytes"
	"encoding/base64"
	"strconv"
)

type FofaInterface interface {
	GetUserInfo(email string, key string)
	GetSearchInfo(query string, page int) (string, int, int)
}

type FofaCore struct {
}

var globalEmail string
var globalKey string

func (c FofaCore) GetUserInfo(email string, key string) {
	globalEmail = email
	globalKey = key
	url := "https://fofa.so/api/v1/info/my?email=" + email + "&key=" + key
	res := httpClient.DoGet(url, nil, nil)
	var userInfo model.FofaUserInfo
	err := json.Unmarshal(res, &userInfo)
	if err != nil {
		log.Log("unmarshal error:"+err.Error(), log.ERROR)
		return
	}
	if userInfo.IsVip != true {
		log.Log("Error API Key", log.ERROR)
		return
	}
	log.Log("Connect Success", log.INFO)
	log.Log("Your Name Is "+userInfo.Username, log.INFO)
	log.Log("Your Email Is "+userInfo.Email, log.INFO)
}

func (c FofaCore) GetSearchInfo(query string, page int) (string, int, int) {
	result := bytes.Buffer{}
	encodeString := base64.StdEncoding.EncodeToString([]byte(query))
	url := "https://fofa.so/api/v1/search/all?email=" +
		globalEmail + "&key=" + globalKey + "&qbase64=" + encodeString +
		"&page=" + strconv.Itoa(page) + "&size=100"
	res := httpClient.DoGet(url, nil, nil)
	var searchInfo model.FofaServiceInfo
	err := json.Unmarshal(res, &searchInfo)
	if err != nil {
		return "", 0, 0
	}
	if searchInfo.Error == true {
		return "", 0, 0
	}
	log.Log("Parsing Data......", log.INFO)
	for i := 0; i < searchInfo.Size; i++ {

	}
	for _, v := range searchInfo.Results {
		result.WriteString(v[1] + ":" + v[2] + "\n")
	}
	return result.String(), searchInfo.Page, searchInfo.Size
}
