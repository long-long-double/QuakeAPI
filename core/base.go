package core

import "QuakeAPI/utils"

var httpClient utils.HttpClient

func init() {
	httpClient = utils.HttpClient{}
}
