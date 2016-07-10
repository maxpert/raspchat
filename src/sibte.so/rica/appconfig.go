package rica

import (
	"io/ioutil"
	"log"

	"encoding/json"
)

type ApplicationConfig struct {
	BindAddress        string            `json:"bind_address"`
	LogFilePath        string            `json:"log_file"`
	DBPath             string            `json:"db_path"`
	AllowHotRestart    bool              `json:"allow_hot_reboot"`
	GCMToken           string            `json:"gcm_token"`
	AllowedOrigins     []string          `json:"allowed_origins"`
	ExternalSignIn     map[string]string `json:"external_sign_in"`
	WebSocketUrl       string            `json:"websocket_url"`
	WebSocketSecureUrl string            `json:"websocketsecure_url"`
}

var CurrentAppConfig ApplicationConfig

func LoadApplicationConfig(filePath string) {
	conf := &CurrentAppConfig
	if filePath == "" {
		conf.AllowHotRestart = false
		conf.BindAddress = ":8080"
		conf.DBPath = "/tmp"
		conf.LogFilePath = ""
		conf.AllowedOrigins = make([]string, 0)
		conf.ExternalSignIn = make(map[string]string)
		conf.WebSocketUrl = ""
		conf.WebSocketSecureUrl = ""
		return
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panic(err)
	}

	if err := json.Unmarshal(content, &CurrentAppConfig); err != nil {
		log.Panic(err)
	}

	log.Println("=== Loaded configuration")
	log.Println(CurrentAppConfig)
}
