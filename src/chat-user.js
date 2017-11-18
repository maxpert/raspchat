const debug = require('debug')('raspchat:chat-user');
const E3 = require('eventemitter3');

const JSONMessage = require('./json-message');
const { buildMessage } = require('./utils');

const DefaultOptions = {
    parser: JSONMessage
};

class ChatUser {
    /**
     *
     * @param {String} id of user
     * @param {Websocket} ws connection
     * @param {String} nick of the user
     * @param {options} options for user
     */
    constructor(id, ws, nick, options) {
        this.events = new E3.EventEmitter();
        this.id = id;
        this.nick = nick;
        this.options = Object.assign({}, DefaultOptions, options);
        this.lastSeen = new Date().getTime();

        this._ws = ws;
        this._joined = new Map();

        debug('User connected with', this.id, this.nick, this.lastSeen);
        this._ws.on('message', message => {
            // Message from user
            this.events.emit(
                'message',
                this.options.parser.fromSerialized(message),
                this);
        });

        this._ws.on('close', () => {
            this.leaveAll();
            this._ws.removeAllListeners('message');
            this._ws.removeAllListeners('close');
            try { this._ws.terminate(); } catch (e) { debug(e); }

            this.events.emit('close', id, ws);
        });

        this.events.on('send', this.send, this);
    }

    /**
     * Gets list of joined rooms
     */
    get joinedRooms() {
        return [...this._joined.values()];
    }

    /**
     * Receive message(s) for user and send it over wire
     * Messages can be JSONMessage(s) or String(s)
     * Each message is sent over socket as seperate payload
     * @param {JSONMessage|JSONMessage[]|String|String[]} messages
     */
    send(messages) {
        if (!Array.isArray(messages)) {
            messages = [messages];
        }

        const shutdownIfError = (err) => {
            if (!err) {
                return;
            }

            debug('Websocket error', err);
            this.events.emit('error', err);
            this._ws.emit('close');
        };

        debug(`Sending user ${this.id} messages`, messages.length);
        for (const msg of messages) {
            try {
                const payload = msg.serialized || msg;
                this._ws.send(
                    payload,
                    { compress: false },
                    shutdownIfError
                );
            } catch (e) {
                shutdownIfError(e);
            }
        }
        debug('Sending payload completed...')
    }

    /**
     * Joins a chat room
     * @param {ChatRoom} chatRoom
     */
    join(chatRoom) {
        debug('User requested to join room', chatRoom.roomName);
        chatRoom.addUser(this.id, this.events);
        this._joined.set(chatRoom.roomName, chatRoom);
        this.events.emit('join', this, [chatRoom]);
        chatRoom.publish(
            buildMessage('group-join', {
                to: chatRoom.roomName,
                from: this.nick
            })
        );

        debug(`User ${this.id} rooms`, this._joined.keys());
    }

    /**
     * Leaves a chat room
     * @param {ChatRoom} chatRoom
     */
    leave(chatRoom) {
        debug('User request to leave room', chatRoom.roomName);
        chatRoom.removeUser(this.id);
        this._joined.delete(chatRoom.roomName);
        this.events.emit('leave', this, [chatRoom]);
        chatRoom.publish(
            buildMessage('group-leave', {
                to: chatRoom.roomName,
                from: this.nick
            })
        );

        debug(`User ${this.id} rooms`, this._joined.keys());
    }

    /**
     * Leave all chat rooms joined by user
     */
    leaveAll() {
        debug('Leaving all rooms for user', this.id);
        [...this._joined.values()].forEach(chatRoom => {
            chatRoom.removeUser(this.id);
            chatRoom.publish(
                buildMessage('group-leave', {
                    to: chatRoom.roomName,
                    from: this.nick
                })
            );
        });

        this._joined.clear();
        this.events.emit('leave', this, [...this._joined]);
        debug(`User ${this.id} rooms`, this._joined.keys());
    }
}

module.exports = ChatUser;
