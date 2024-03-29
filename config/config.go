package config

import (
	"log"
	"todo_app/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port string
	SQLDriver string
	DbName string
	LogFile string
	Static string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSetting(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").MustString("sqlite3"),
		DbName: cfg.Section("db").Key("name").String(),
		LogFile: cfg.Section("web").Key("logfile").String(),
		// 静的ファイルがありパスをconfig.iniで設定
		Static: cfg.Section("web").Key("static").String(),
	}
}