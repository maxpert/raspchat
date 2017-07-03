#!/bin/bash
export GOPATH=`pwd`

########### Go vend
echo "Installing govend"

go get github.com/govend/govend
########### Packages
echo "Installing Go packages"
pushd src
../bin/govend -v
popd

########### NPM
# echo "Installing NPM packages"
# npm install
