const http = require('http');
const debug = require('debug')('raspchat:upgrade-middleware');

let wss = null;

module.exports = function (app, WebSocketServer) {
    debug('Upgrade middleware...');
    if (!WebSocketServer || !app) {
        throw new Error('Websocket upgrade middleware requires Express app and WebSocket object');
    }

    wss = wss || new WebSocketServer({
        noServer: true
    });

    wss.on('error', console.error);
    wss.on('headers', h => debug(h));

    return function (req, socket, upgradeHead) {
        debug('Upgrading middleware for req', req.url);
        var res = new http.ServerResponse(req);
        res.assignSocket(socket);

        // avoid hanging onto upgradeHead as this will keep the entire
        // slab buffer used by node alive
        var head = new Buffer(upgradeHead.length);
        upgradeHead.copy(head);

        res.on('finish', function () {
            debug('Destroying response socket');
            res.socket.destroy();
        });

        res.websocket = function (cb) {
            wss.handleUpgrade(req, socket, head, function (client) {
                wss.emit('connection', client, req);
                cb && cb(client);
            });
        };

        return app(req, res);
    };
};
