const IntFormat = require('biguint-format');
const FlakeId = require('flake-idgen');
const ShortId = require('shortid');

const JSONMessage = require('./json-message');

const IdGen = new FlakeId();
const DefaultGenIdOptions = {
    short: false,
    encoding: 'hex'
};

function genId(opts) {
    opts = Object.assign({}, DefaultGenIdOptions, opts);
    if (opts.short) {
        return ShortId();
    }

    return IntFormat(IdGen.next(), opts.encoding);
}

function buildMessage(type, extra) {
    return JSONMessage.fromObject(
        Object.assign({
            '@': type,
            '!id': genId(),
            'utc_timestamp': new Date().getTime()
        }, extra)
    );
}

module.exports = {
    genId,
    buildMessage
};
