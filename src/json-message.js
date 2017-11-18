function _DeepFreeze(obj) {
    Object.getOwnPropertyNames(obj).forEach(function(name) {
        var prop = obj[name];
        if (typeof prop == 'object' && prop !== null) {
            _DeepFreeze(prop);
        }
    });

    return Object.freeze(obj);
}

class JSONMessage {
    constructor(parsed, serialized) {
        this.parsed = _DeepFreeze(parsed);
        this.serialized = serialized;
    }
};

JSONMessage.fromSerialized = function (str) {
    const obj = JSON.parse(str);
    return new JSONMessage(obj, str);
};

JSONMessage.fromObject = function (obj) {
    const str = JSON.stringify(obj);
    return new JSONMessage(obj, str);
};

module.exports = JSONMessage;
