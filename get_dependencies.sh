#!/bin/bash

echo "Installing Go packages"
curl -s https://raw.githubusercontent.com/pote/gpm/v1.4.0/bin/gpm > gpm
chmod +x gpm
./gpm get

echo "Installing NPM packages"
npm install