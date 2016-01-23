#!/bin/bash

env GOPATH=`pwd` go get github.com/speps/go-hashids
env GOPATH=`pwd` go get github.com/gorilla/websocket
env GOPATH=`pwd` go get github.com/julienschmidt/httprouter
env GOPATH=`pwd` go get github.com/boltdb/bolt/...
env GOPATH=`pwd` go get gopkg.in/natefinch/lumberjack.v2
env GOPATH=`pwd` go get github.com/fvbock/endless
env GOPATH=`pwd` go get github.com/googollee/go-gcm


npm install bower
npm install uglify-js
