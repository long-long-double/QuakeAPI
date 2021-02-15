package model

type QuakeUserInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID   string `json:"id"`
		User struct {
			ID       string `json:"id"`
			Username string `json:"username"`
			Fullname string `json:"fullname"`
			Email    string `json:"email"`
		} `json:"user"`
		Baned            bool   `json:"baned"`
		BanStatus        string `json:"ban_status"`
		Credit           int    `json:"credit"`
		PersistentCredit int    `json:"persistent_credit"`
		AvatarID         string `json:"avatar_id"`
		Token            string `json:"token"`
		MobilePhone      string `json:"mobile_phone"`
		Source           string `json:"source"`
		PrivacyLog       struct {
			Status bool   `json:"status"`
			Time   string `json:"time"`
		} `json:"privacy_log"`
		EnterpriseInformation struct {
			Name   interface{} `json:"name"`
			Email  interface{} `json:"email"`
			Status string      `json:"status"`
		} `json:"enterprise_information"`
		PersonalInformationStatus bool `json:"personal_information_status"`
		Role                      []struct {
			Fullname string `json:"fullname"`
			Priority int    `json:"priority"`
			Credit   int    `json:"credit"`
		} `json:"role"`
	} `json:"data"`
	Meta struct {
	} `json:"meta"`
}

type FofaUserInfo struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	Fcoin      int    `json:"fcoin"`
	IsVip      bool   `json:"isvip"`
	VipLevel   int    `json:"vip_level"`
	IsVerified bool   `json:"is_verified"`
	Avatar     string `json:"avatar"`
	Message    int    `json:"message"`
	FofaCliVer string `json:"fofacli_ver"`
	FofaServer bool   `json:"fofa_server"`
}
