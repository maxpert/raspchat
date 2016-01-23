package main

/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/fvbock/endless"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/natefinch/lumberjack.v2"

	"sibte.so/rica"
)

var index_bytes []byte

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if index_bytes == nil {
		filedat, err := ioutil.ReadFile("static/index.html")

		if err != nil {
			fmt.Fprintf(w, "hit /static/test.html")
		}

		index_bytes = filedat
	}

	w.Write(index_bytes)
}

func clearCache(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	index_bytes = nil

	fmt.Fprint(w, "Done")
}

var _groupInfoManager rica.GroupInfoManager = rica.NewInMemoryGroupInfo()
var _nickRegistry *rica.NickRegistry = rica.NewNickRegistry()

func _installSocketMux(mux *http.ServeMux) (err error) {
	err = nil

	if err != nil {
		log.Fatal(err)
		return
	}

	s := rica.NewChatService().WithRESTRoutes("/chat")

	mux.Handle("/chat", s)
	mux.Handle("/chat/", s)
	return
}

func _installHttpRoutes(mux *http.ServeMux) (err error) {
	err = nil
	router := httprouter.New()
	router.GET("/", index)
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

	_installSocketMux(mux)
	_installHttpRoutes(mux)

	endless.DefaultHammerTime = 10 * time.Second
	if conf.AllowHotRestart == false {
		server := &http.Server{
			Addr:    conf.BindAddress,
			Handler: mux,
		}

		log.Println("Starting server...", conf.BindAddress)
		log.Panic(server.ListenAndServe())
	} else {
		endless.ListenAndServe(conf.BindAddress, mux)
	}
}
