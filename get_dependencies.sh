#!/bin/bash

echo "Installing Go packages"
env GOPATH=`pwd` go get github.com/Azure/go-autorest/autorest
env GOPATH=`pwd` go get golang.org/x/net/context
env GOPATH=`pwd` go get golang.org/x/text
env GOPATH=`pwd` go get github.com/speps/go-hashids
env GOPATH=`pwd` go get github.com/gorilla/websocket
env GOPATH=`pwd` go get github.com/julienschmidt/httprouter
env GOPATH=`pwd` go get github.com/boltdb/bolt/...
env GOPATH=`pwd` go get gopkg.in/natefinch/lumberjack.v2
env GOPATH=`pwd` go get github.com/googollee/go-gcm
env GOPATH=`pwd` go get github.com/Azure/azure-sdk-for-go/management

pushd src/github.com/speps/go-hashids
git checkout -q master
git checkout -q tags/v1.0.0
popd > /dev/null

pushd src/github.com/gorilla/websocket
git checkout -q master
git checkout -q tags/v1.0.0
popd > /dev/null

pushd src/github.com/julienschmidt/httprouter
git checkout -q master
git checkout -q tags/v1.1
popd > /dev/null

pushd src/github.com/boltdb/bolt
git checkout -q master
git checkout -q tags/v1.2.1
popd > /dev/null

echo "Installing NPM packages"
npm install
