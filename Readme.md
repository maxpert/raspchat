# An IRC like chat system written in Node.js

[![Build Status](https://semaphoreci.com/api/v1/projects/bb232bf4-f866-4d30-a526-218e9f3aed5b/647934/badge.svg)](https://semaphoreci.com/maxpert/raspchat)

# Demo

Basic demo is available [Here](http://raspchat.com).

## Why?
I had a spare Raspberry Pi and I wanted to use it! One of ideas in my head was to have your own on premis chat server that you can use for cheap and own your data (< $50 hardware) forever and free!

## Requirements
For compiling you need:
 * Node 8+ (with npm5)
 * libsqlite3
 * OpenSSL 1.x.x (for uWebsockets)
 * zlib 1.x (for uWebsockets)
 * libuv 1.3+ or Boost.Asio 1.x (both optional on Linux)

You can use following commands to install them:
 * Fedora: `sudo dnf install openssl-devel zlib-devel`
 * Homebrew: `brew install openssl zlib libuv`
 * Vcpkg: `vcpkg install openssl zlib libuv and/or vcpkg install openssl:x64-windows zlib:x64-windows libuv:x64-windows`

Once you have installed dependencies above just do `npm install && gulp` (creates a dist folder that you can upload to your machine). Project can run on almost any machine that nodejs supports. I have successfully tested it on Raspberry Pi, Orange Pi etc.

## Configurations

You can use a `.env` file or environment variables to configure the server, here are the environment variables you can change:

 * `RC_URL` (default `http://localhost:3000/`) **R**asp**c**hat **URL** where server would start listening
 * `DB_PATH` (default `chat-log.db`) **D**ata**b**ase **path** for SQLite database
 * `WS_URL` (default `ws://localhost:3000/chat`) **W**eb**s**ocket **URL** if you are running it behind some proxy
 * `WSS_URL` (default `wss://localhost:3000/chat`) **W**eb**s**ocket **s**ecure **URL** if you are running it behind some proxy

You can read about `.env` file [here](https://www.npmjs.com/package/dotenv) 

## Features:

 * Basic GIF support
 * Basic nick support
 * Channel (join/leave/list-members) support
 * Markdown support
 * Message history on group join

## Pending:

 * File upload support
 * Refactor frontend to Vue2 and cleanup frontend
 * Unit tests
 * Introduce admin panel with:
   * Reserved alias authorization
   * IP limiting/banning
   * Channel management and permissions
