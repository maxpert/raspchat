const debug = require('debug')('raspchat:chat-room');

const { ChatLog } = require('./storage');
const $meta = Symbol('meta');

class ChatRoom {
    constructor(name, enableLog) {
        this[$meta] = {
            name,
            enableLog,
            users: new Map(),
            publishQueue: {
                immediate: null,
                msgs: []
            }
        };

        debug('Room created', name, 'with logging', enableLog);
    }

    /**
     * Get name of room
     */
    get roomName() {
        return this[$meta].name;
    }

    /**
     * Returns true if room has users;
     * false otherwise
     */
    get hasUsers() {
        return this[$meta].users.size !== 0;
    }

    /**
    * Returns list of all the user ids currently
    * member of the room
    */
    get userIds() {
        return [...this[$meta].users.keys()];
    }

    /**
     * Add user to room
     * @param {string} id of user
     * @param {EventEmitter} eventEmitter that can receive messages
     */
    addUser(id, eventEmitter) {
        this[$meta].users.set(id, eventEmitter);
        return this;
    }

    /**
     * Remove user from room
     * @param {string} id of user
     */
    removeUser(id) {
        this[$meta].users.delete(id);
        return this;
    }

    /**
     * Publish message to all users
     * @param {any} msg to publish
     * @param {any} options for publish
     *              forceFlush (false) forces publish flushing right away
     */
    publish(msg, options) {
        const queue = this[$meta].publishQueue;
        const users = this[$meta].users;

        const flushMessages = function () {
            const allMessages = queue.msgs;
            queue.msgs = [];
            queue.immediate = null;

            for (const [id ,eventEmitter] of users) {
                eventEmitter.emit('send', allMessages);
            }
        };

        // Merge options
        options = Object.assign({}, {
            forceFlush: false,
            limitFlush: 100,
            flushDelay: 10
        }, options);

        queue.msgs.push(msg);

        // Record message
        if (this[$meta].enableLog && msg.parsed && msg.parsed['to'] === this.roomName) {
            debug('Saving chat log', msg.serialized);
            ChatLog.put(msg.parsed['!id'], msg.parsed['to'], msg.serialized).catch(debug);
        }

        // Clear previous timeout
        if (queue.immediate) {
            clearTimeout(queue.immediate);
        }

        // Forced flush
        if (options.forceFlush || queue.msgs.length > options.limitFlush) {
            flushMessages();
            return;
        }

        queue.immediate = setTimeout(flushMessages, options.flushDelay);
    }
}

module.exports = ChatRoom;
