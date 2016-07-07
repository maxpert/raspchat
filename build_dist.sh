#!/bin/bash

rm -rf dist
mkdir -p dist
mkdir -p dist/static

echo "Compiling server..."
./build_server.sh

echo "Packaging client assets..."
./compile_assets.sh
rm -rf dist/static
cp -r ./static ./dist/static
rm -rf dist/static/bower_components
rm dist/static/chat.dev.html
rm -rf dist/static/components
rm dist/static/*.mp3
rm dist/static/file_transfer.js dist/static/core.js dist/static/peer_negotiator.js dist/static/rtc.js
