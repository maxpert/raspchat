package main

/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

import (
    "flag"
    "log"
    "net/http"

    "github.com/julienschmidt/httprouter"
    "gopkg.in/natefinch/lumberjack.v2"

    "sibte.so/rasconfig"
    "sibte.so/rasweb"
    "sibte.so/rica"
)

func installSocketMux(mux *http.ServeMux, appConfig rasconfig.ApplicationConfig) (err error) {
    err = nil
    s := rica.NewChatService(appConfig).WithRESTRoutes("/chat")

    mux.Handle("/chat", s)
    mux.Handle("/chat/", s)
    return
}

var routeHandlers = []rasweb.RouteHandler{
    rasweb.NewGifHandler(),
    rasweb.NewFileUploadHandler(),
    rasweb.NewConfigRouteHandler(),
    rasweb.NewDirectPagesHandler(),
}

func installHTTPRoutes(mux *http.ServeMux) (err error) {
    err = nil
    router := httprouter.New()

    // Register all routes
    for _, h := range routeHandlers {
        if err := h.Register(router); err != nil {
            log.Panic("Unable to register route")
        }
    }

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
    rasconfig.LoadApplicationConfig(parseArgs())
    conf := rasconfig.CurrentAppConfig

    if conf.LogFilePath != "" {
        log.SetOutput(&lumberjack.Logger{
            Filename:   conf.LogFilePath,
            MaxBackups: 3,
            MaxSize:    5,
            MaxAge:     15,
        })
    }

    mux := http.NewServeMux()
    installSocketMux(mux, conf)
    installHTTPRoutes(mux)
    server := &http.Server{
        Addr:    conf.BindAddress,
        Handler: mux,
    }

    log.Println("Starting server...", conf.BindAddress)
    log.Panic(server.ListenAndServe())
}
