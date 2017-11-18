const FlakeId = require('flake-idgen');
const Inflector = require('inflected');
const IntFormat = require('biguint-format');

const ChatUser = require('./chat-user');
const ChatRoom = require('./chat-room');
const JSONMessage = require('./json-message');
const CommandHandlers = require('./command-handlers');
const NickRegistery = require('./nick-registery');

const DefaultConfig = {
    commandRoom: 'SERVER'
};

const Rooms = new Map();
const Users = new Map();
const IdGen = new FlakeId();

function genId() {
    return IntFormat(IdGen.next(), 'hex');
}

function GetRoom(maybeName) {
    const name = maybeName || DefaultConfig.commandRoom;
    const chatRoom = Rooms.has(name) ? Rooms.get(name) : new ChatRoom(name, name !== DefaultConfig.commandRoom);
    Rooms.set(name, chatRoom);
    return chatRoom;
}

function UserDisconnected(id) {
    if (Users.has(id)) {
        const user = Users.get(id);
        user.events.removeAllListeners('message');
        user.events.removeAllListeners('close');
        NickRegistery.delete(id);
        Users.delete(id);
    }
}

function UserMessage(message, user) {
    const command = message.parsed['@'].replace('-', '_');
    const methodName = Inflector.camelize(command, false);
    const commandMethod = CommandHandlers[methodName];
    if (!commandMethod) {
        console.error('Invalid command', command, methodName);
        return;
    }

    commandMethod(user, GetRoom, message.parsed).catch(console.error);
}

function ServerPinger() {
    const serverRoom = GetRoom();
    serverRoom.publish(JSONMessage.fromObject({
        '@': 'ping',
        '!id': genId(),
        'utc_timestamp': 0,
        't': new Date().getTime()
    }));
    setTimeout(ServerPinger, 10000 + Math.random() * 15000);
}

function SendWelcome(user) {
    const serverRoom = GetRoom();
    user.send(JSONMessage.fromObject({
        '@': serverRoom.roomName,
        '!id': genId(),
        'utc_timestamp': new Date().getTime(),
        'msg': 'Welcome...'
    }));
}

module.exports = function () {
    ServerPinger();

    return function (req, res) {
        if (!res.websocket) {
            res.status(500).send('Not a websocket connection');
            return;
        }

        let id = genId();
        while (Users.has(id)) {
            id = genId();
        }

        res.websocket(function (ws) {
            const user = new ChatUser(id, ws, NickRegistery.create(id));

            SendWelcome(user);
            user.events.once('close', UserDisconnected);
            user.events.on('message', UserMessage);

            // Make user join SERVER
            user.join(GetRoom());
            Users.set(id, user);
        });
    };
};
