#!/bin/bash

echo "Installing Go packages"
curl -s https://raw.githubusercontent.com/pote/gpm/v1.4.0/bin/gpm > gpm
chmod +x gpm
env GOPATH=`pwd` ./gpm get
