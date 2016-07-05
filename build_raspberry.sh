#!/bin/bash

env GOPATH=`pwd` GOOS=linux GOARCH=arm GOARM=6 go build -o arm-server sibte.so
env GOPATH=`pwd` GOOS=linux GOARCH=386 go build -o chat-server-32 sibte.so
env GOPATH=`pwd` GOOS=linux go build -o chat-server sibte.so
