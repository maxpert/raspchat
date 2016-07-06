#!/bin/bash

echo "Compiling Linux ARM6"
env GOPATH=`pwd` GOOS=linux GOARCH=arm GOARM=6 go build -o arm-server sibte.so

echo "Compiling Linux x86"
env GOPATH=`pwd` GOOS=linux GOARCH=386 go build -o chat-server-32 sibte.so

echo "Compiling Linux x64"
env GOPATH=`pwd` GOOS=linux go build -o chat-server sibte.so

mv arm-server ./dist
mv chat-server-32 ./dist
mv chat-server ./dist
