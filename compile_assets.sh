#!/bin/bash

export PATH=`pwd`/node_modules/.bin:$PATH

pushd static
bower install vue
bower install markdown-it
bower install moment
bower install qwest
bower install he
popd

uglifyjs --compress --mangle --output static/app.min.js -- \
  static/bower_components/vue/dist/vue.js \
  static/bower_components/moment/moment.js \
  static/bower_components/qwest/qwest.min.js \
  static/bower_components/markdown-it/dist/markdown-it.js \
  static/bower_components/he/he.js \
  static/core.js static/rtc.js static/peer_negotiator.js \
  static/file_transfer.js \
  static/components/filters.js \
  static/components/app-bar.js \
  static/components/chat-compose.js \
  static/components/chat-log.js \
  static/components/chat-message.js \
  static/components/chrome-bar.js \
  static/components/group-list.js \
  static/components/toggle-buttons.js \
  static/components/sign-in.js \
  static/components/chat.js
