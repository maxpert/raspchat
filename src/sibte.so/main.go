package main

/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/natefinch/lumberjack.v2"

	"sibte.so/rica"
)

var indexPageCache []byte

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if indexPageCache == nil {
		filedat, err := ioutil.ReadFile("static/index.html")

		if err != nil {
			log.Println("HIT /static/test.html")
		}

		indexPageCache = filedat
	}

	w.Write(indexPageCache)
}

func clearCache(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	indexPageCache = nil
	fmt.Fprint(w, "Done")
}

func getChatConfig(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	appConfig := rica.CurrentAppConfig
	isJs := false

	if strings.HasSuffix(params.ByName("type"), ".js") {
		isJs = true
	}

	if isJs {
		w.Header().Add("Content-Type", "text/javascript")
	} else {
		w.Header().Add("Content-Type", "application/json")
	}

	config := make(map[string]interface{})
	config["webSocketConnectionUri"] = appConfig.WebSocketUrl
	config["webSocketSecureConnectionUri"] = appConfig.WebSocketSecureUrl
	config["externalSignIn"] = appConfig.ExternalSignIn

	if isJs {
		fmt.Fprint(w, "window.RaspConfig=")
	}

	json.NewEncoder(w).Encode(config)
}

var _groupInfoManager = rica.NewInMemoryGroupInfo()
var _nickRegistry = rica.NewNickRegistry()

func _installSocketMux(mux *http.ServeMux, appConfig *rica.ApplicationConfig) (err error) {
	err = nil

	if err != nil {
		log.Fatal(err)
		return
	}

	s := rica.NewChatService(appConfig).WithRESTRoutes("/chat")

	mux.Handle("/chat", s)
	mux.Handle("/chat/", s)
	return
}

func _installHTTPRoutes(mux *http.ServeMux) (err error) {
	err = nil
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/config/:type", getChatConfig)
	router.GET("/gif", rica.FindRightGif)
	router.GET("/_clear", clearCache)
	router.ServeFiles("/static/*filepath", http.Dir("./static"))

	mux.Handle("/", router)
	return
}

func parseArgs() (filePath string) {
	flag.StringVar(&filePath, "config", "", "Path to configuration file")
	flag.Parse()
	return
}

func main() {
	rica.LoadApplicationConfig(parseArgs())
	conf := rica.CurrentAppConfig

	if conf.LogFilePath != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   conf.LogFilePath,
			MaxBackups: 3,
			MaxSize:    5,
			MaxAge:     15,
		})
	}

	mux := http.NewServeMux()

	_installSocketMux(mux, &conf)
	_installHTTPRoutes(mux)

	server := &http.Server{
		Addr:    conf.BindAddress,
		Handler: mux,
	}

	log.Println("Starting server...", conf.BindAddress)
	log.Panic(server.ListenAndServe())
}
