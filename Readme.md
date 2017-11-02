# A basic Websocket chat server in Go

[![Build Status](https://semaphoreci.com/api/v1/projects/bb232bf4-f866-4d30-a526-218e9f3aed5b/647934/badge.svg)](https://semaphoreci.com/maxpert/raspchat)

## Why?
I had a spare Raspberry Pi and I wanted to use it! One of ideas in my head was to have your own on premis chat server that you can use for cheap and own your data (< $50 hardware) forever and free! So I took a look around and found golang was perfect fit, and I made one for everyone :)

## Requirements
For compiling you need:
 * Go 1.5+
 * Node 8+ (with npm)

Once you have installed both just cd to directory and run ```./get_dependencies.sh && ./build_dist.sh``` (creates a dist folder). Project can run on almost any machine that go can cross compile to. I have successfully tested it on Raspberry Pi, Orange Pi, even Samsung Galaxy S1 mobile phone.

# Demo

Basic demo is available [Here](http://raspchat.com).

## Features:

 * Basic GIF support
 * Basic nick support
 * Channel support
 * Markdown support
 * Message history support
 * File upload support

## Pending:

 * Improve build and deploy script
 * Introduce admin panel with:
   * Reserved alias authorization
   * IP limiting/banning
   * Channel management and permissions
