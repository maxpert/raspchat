#!/bin/bash

rm -rf dist
mkdir -p dist
mkdir -p dist/static

echo "Compiling server..."
./build_server.sh

echo "Packaging client assets..."
./compile_assets.sh

echo "Generating download package..."
tar -zcvf /tmp/raspchat-current.tar.gz ./dist
mv /tmp/raspchat-current.tar.gz ./dist/static

echo "----- Fin."