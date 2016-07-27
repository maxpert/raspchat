#!/bin/bash

echo "Compiling Linux ARM6"
env GOPATH=`pwd` GOOS=linux GOARCH=arm GOARM=6 go build -o arm-server sibte.so

echo "Compiling Linux 32-bit"
env GOPATH=`pwd` GOOS=linux GOARCH=386 go build -o chat-server-32 sibte.so

echo "Compiling Linux 64-bit"
env GOPATH=`pwd` GOOS=linux GOARCH=amd64 go build -o chat-server sibte.so

echo "Compiling MacOS 32-bit"
env GOPATH=`pwd` GOOS=darwin GOARCH=386 go build -o macos-chat-server-32 sibte.so

echo "Compiling MacOS 64-bit"
env GOPATH=`pwd` GOOS=darwin GOARCH=amd64 go build -o macos-chat-server sibte.so

echo "Compiling Windows 32-bit"
env GOPATH=`pwd` GOOS=windows GOARCH=386 go build -o windows-chat-server-32.exe sibte.so

echo "Compiling Windows 64-bit"
env GOPATH=`pwd` GOOS=windows GOARCH=amd64 go build -o windows-chat-server.exe sibte.so

mv arm-server ./dist
mv chat-server-32 ./dist
mv chat-server ./dist
mv macos-chat-server-32 ./dist
mv macos-chat-server ./dist
mv windows-chat-server-32.exe ./dist
mv windows-chat-server.exe ./dist
