package core

import (
	"QuakeAPI/utils"
	jsoniter "github.com/json-iterator/go"
)

var httpClient utils.HttpClient
var json jsoniter.API

func init() {
	httpClient = utils.HttpClient{}
	json = jsoniter.ConfigCompatibleWithStandardLibrary
}
