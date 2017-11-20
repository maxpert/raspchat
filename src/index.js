require('dotenv').config();

const debug = require('debug')('raspchat:index');
const express = require('express');
const WebSocket = require('uws');
const http = require('http');
const url = require('url');

const upgradeHandler = require('./upgrade-middleware');
const chatHandler = require('./chat-handler');
const chatAPI = require('./api');

const app = express();
const server = http.createServer(app);
const rcUrl = url.parse(process.env.RC_URL || 'http://localhost:3030/');

chatAPI(app, rcUrl);
app.use('/chat', chatHandler());
app.use('/static', express.static('static'));
app.get('/', (req, res) => res.redirect(301, '/static'));

server.on('upgrade', upgradeHandler(app, WebSocket.Server));
server.listen(~~rcUrl.port, rcUrl.hostname, function () {
    debug('Listening on ...', server.address().port);
});
