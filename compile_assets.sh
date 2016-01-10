#!/bin/bash

export PATH=`pwd`/node_modules/.bin:$PATH

pushd static
bower install vue
bower install markdown-it
bower install moment
popd

uglifyjs --compress --mangle --output static/app.min.js -- static/bower_components/vue/dist/vue.js static/bower_components/moment/moment.js static/bower_components/markdown-it/dist/markdown-it.js static/core.js static/rtc.js static/peer_negotiator.js static/file_transfer.js static/chat.js
