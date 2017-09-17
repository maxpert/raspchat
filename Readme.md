# A basic Websocket chat server in Go

[![Build Status](https://semaphoreci.com/api/v1/projects/bb232bf4-f866-4d30-a526-218e9f3aed5b/647934/badge.svg)](https://semaphoreci.com/maxpert/raspchat)

## Requirements
For compiling you need:
 * Go 1.5+
 * Node 5+ (with npm)

Once you have installed both just cd to directory and run ```./get_dependencies.sh && ./build_dist.sh``` (creates a dist folder). Project can run on almost any machine that go can cross compile to.

# Demo

Basic demo is available [Here](http://beta.raspchat.com).

## Features:

 * Basic GIF support
 * Basic nick support
 * Channel support
 * Markdown support
 * Message history support
 * File upload support
 * GCM push notification support (incomplete)


## Pending:

 * Improve build and deploy script
 * Introduce admin panel with:
   * Reserved alias authorization
   * IP limiting/banning
   * Channel management and permissions
