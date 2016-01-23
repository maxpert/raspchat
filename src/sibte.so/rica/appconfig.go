package rica

import (
	"io/ioutil"
	"log"

	"encoding/json"
)

type ApplicationConfig struct {
	BindAddress     string `json:"bind_address"`
	LogFilePath     string `json:"log_file"`
	DBPath          string `json:"db_path"`
	AllowHotRestart bool   `json:"allow_hot_reboot"`
	GCMToken        string `json:"gcm_token"`
}

var CurrentAppConfig ApplicationConfig

func LoadApplicationConfig(filePath string) {
	conf := &CurrentAppConfig
	if filePath == "" {
		conf.AllowHotRestart = false
		conf.BindAddress = ":8080"
		conf.DBPath = "/tmp"
		conf.LogFilePath = ""
		return
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panic(err)
	}

	if err := json.Unmarshal(content, &CurrentAppConfig); err != nil {
		log.Panic(err)
	}
}
