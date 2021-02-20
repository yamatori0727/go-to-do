package config

import (
	"awesomeProject/awesomeProject/todo_app/utils"
	"log"
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct{
	Port string
	SQLDriver string
	DbName string
	LogFile string
}

var Config ConfigList

func init(){
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

//プログラムの中身が動くための詳細設定
func LoadConfig(){
	cfg, err := ini.Load("config.ini")
	if err != nil{
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName: cfg.Section("db").Key("name").String(),
		LogFile: cfg.Section("web").Key("logfile").String(),
	}

}
