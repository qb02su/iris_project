package config

import (
	"encoding/json"
	"os"
)

//服务端配置
type AppConfig struct {
	AppName    string   `json:"app_name"`
	Port       string   `json:"port"`
	StaticPath string   `json:"static_path"`
	Mode       string   `json:"mode"`
	DataBase   DataBase `json:"data_base"`
	Redis      Redis    `json:"redis"`
}

/**
 * mysql配置
 */
type DataBase struct {
	Drive    string `json:"drive"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Host     string `json:"host"`
	Database string `json:"database"`
}

/**
 * Redis 配置
 */
type Redis struct {
	NetWork  string `json:"net_work"`
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Prefix   string `json:"prefix"`
}

//初始化服务器配置
func InitConfig() *AppConfig {
	var config *AppConfig
	file, err := os.Open("D:/iriswork/config.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err.Error())
	}
	//config = &AppConfig{}
	return config
}
