const bluebird = require('bluebird');
const debug = require('debug')('raspchat:api');
const xhr = bluebird.promisifyAll(require('request'));

const { ChatLog, GifCache } = require('./storage');

async function findGif(keywords) {
    const gifResp = await xhr.postAsync('https://rightgif.com/search/web', {
        form: {
            text: keywords
        }
    });

    return JSON.parse(gifResp.body);
}

module.exports = function (app, rcUrl) {
    app.get('/api/chat/channel/:channel/messages', async (req, res, next) => {
        const limit = (~~req.query.limit) || 50;
        const offset = (~~req.query.offset) || 0;
        const start_id = req.query.start_id || '';

        try {
            const tuples = await ChatLog.fetchFor(req.params.channel, limit, start_id);
            const messages = tuples.map(t => JSON.parse(t.message));

            res.send({
                id: req.params.channel,
                limit,
                offset,
                start_id,
                messages
            });
        } catch (e) {
            debug('Error while getting channel messages', e);
            next(e);
        }
    });

    app.get('/gif', async (req, res, next) => {
        const keywords = (req.query.q || 'void').toLowerCase();
        try {
            const cachedUrls = await GifCache.get(keywords);
            debug('Cached resultset', cachedUrls);
            if (cachedUrls && cachedUrls.length >= 1) {
                debug('Cache hit for gif', keywords);
                res.send({
                    url: cachedUrls[0].url
                });
                return;
            }

            const gifObj = await findGif(keywords);
            const putPromise = GifCache.put(keywords, gifObj.url);
            res.send({
                url: gifObj.url
            });

            await putPromise;
            debug('Gif cached...');
        } catch (e) {
            debug('Error while finding gif', e);
            next(e);
        }
    });

    app.get('/config/client.js', function (req, res) {
        const config = {
            'externalSignIn': process.env.EXTERNAL_SIGIN_IN || null,
            'hasAuthProviders': process.env.HAS_AUTH_PROVIDERS || false,
            'webSocketConnectionUri': process.env.WS_URL || `ws://${rcUrl.host}/chat`,
            'webSocketSecureConnectionUri': process.env.WSS_URL || `wss://${rcUrl.host}/chat`
        };

        res.setHeader('Content-Type', 'application/javascript');
        res.send('window.RaspConfig=' + JSON.stringify(config) + ';');
    });
};
