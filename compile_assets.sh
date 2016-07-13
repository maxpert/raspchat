#!/bin/bash

export PATH=`pwd`/node_modules/.bin:$PATH

bower install
gulp compile-assets