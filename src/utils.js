const IntFormat = require('biguint-format');
const FlakeId = require('flake-idgen');

const JSONMessage = require('./json-message');

const IdGen = new FlakeId();

function genId() {
    return IntFormat(IdGen.next(), 'hex');
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
