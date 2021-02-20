package model

type YamlConfig struct {
	Login struct {
		Email    string `yaml:"email"`
		Key      string `yaml:"key"`
		Userinfo bool   `yaml:"userinfo"`
	} `yaml:"login"`
	Search struct {
		Query  string `yaml:"query"`
		Output string `yaml:"output"`
		Total  int    `yaml:"total"`
	} `yaml:"search"`
	Use struct {
		Quake bool `yaml:"quake"`
		Fofa  bool `yaml:"fofa"`
	} `yaml:"use"`
	MySQL struct {
		Use      bool   `yaml:"use"`
		Server   string `yaml:"server"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mysql"`
}
