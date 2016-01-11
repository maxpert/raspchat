# A basic Websocket chat server in Go

[![Build Status](https://semaphoreci.com/api/v1/projects/bb232bf4-f866-4d30-a526-218e9f3aed5b/647934/badge.svg)](https://semaphoreci.com/maxpert/raspchat)

## Requirements
For compiling you need:
 * Go 1.5+
 * Node 5+ (with npm)

Once you have installed both just cd to directory and run ```./get_dependencies.sh && ./build_dist.sh``` (creates a dist folder). Project can run on vitally any machine go can cross compile to.

# Demo

Basic demo is available [Here](http://raspchat.sibte.so). Do note the server is running on Raspberry Pi 1 Model B.

## Features:

 * Basic GIF support
 * Basic nick support
 * Channel support
 * Markdown support


## TODO:

 * Improve build and deploy script
 * Introduce admin panel with:
   * Reserved alias authorization
   * IP limiting/banning
   * Channel management and permissions
 * Loadable extension system
 * Scheduled chat export system (TBD)
