package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

type ConfigList struct {
	ApiKey  string
	LogFile string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}
	Config = ConfigList{
		ApiKey:  cfg.Section("yahoo_finance").Key("api_key").String(),
		LogFile: cfg.Section("sys_trading").Key("log_file").String(),
	}
}
