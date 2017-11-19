const Inflector = require('inflected');
const bluebird = require('bluebird');
const Figlet = bluebird.promisifyAll(require('figlet'));

const ChatUser = require('./chat-user');
const ChatRoom = require('./chat-room');
const JSONMessage = require('./json-message');
const CommandHandlers = require('./command-handlers');
const NickRegistery = require('./nick-registery');
const {genId} = require('./utils');

const DefaultConfig = {
    commandRoom: 'SERVER'
};

const Rooms = new Map();
const Users = new Map();

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
    setTimeout(ServerPinger, 5000 + Math.random() * 20000);
}

async function SendWelcome(user) {
    const serverRoom = GetRoom();
    const normalized_nick = user.nick.replace(/[^\d\w]/, ' ');
    const message = '```\n'+ await Figlet.textAsync('Hi ' + normalized_nick, {
        font: 'The Edge'
    }) + '\n```';
    
    user.send(JSONMessage.fromObject({
        '@': serverRoom.roomName,
        '!id': genId(),
        'utc_timestamp': new Date().getTime(),
        'msg': message
    }));
}

module.exports = function () {
    ServerPinger();

    return function (req, res) {
        if (!res.websocket) {
            res.status(500).send('Not a websocket connection');
            return;
        }

        let id = genId({short: true});
        while (Users.has(id)) {
            id = genId({short: true});
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
