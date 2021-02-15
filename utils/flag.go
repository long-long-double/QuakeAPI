package utils

import (
	"QuakeAPI/model"
	"flag"
)

func ParseInput() model.Input {
	var userInfo bool
	var key string
	var search string
	var help bool
	var output string
	var total int
	var fofa bool
	var quake bool
	var email string
	result := model.Input{}
	flag.StringVar(&key, "key", "", "Input Your API Key.")
	flag.StringVar(&email, "email", "", "If You Use Fofa,You Should Enter This.")
	flag.IntVar(&total, "total", 100, "Number Of Queries You Want.")
	flag.StringVar(&search, "search", "", "Input Search String.")
	flag.StringVar(&output, "output", "result.txt", "Output File.")
	flag.BoolVar(&userInfo, "userinfo", false, "Show Your User Information.")
	flag.BoolVar(&help, "help", false, "Show Help Information.")
	flag.BoolVar(&fofa, "fofa", false, "Use Fofa.")
	flag.BoolVar(&quake, "quake", false, "Use Quake.")
	flag.Parse()
	if key == "" || help == true {
		flag.PrintDefaults()
		return result
	}
	result.UserInfo = userInfo
	result.Key = key
	result.Search = search
	result.Output = output
	result.Total = total
	result.Fofa = fofa
	result.Quake = quake
	result.Email = email
	return result
}
