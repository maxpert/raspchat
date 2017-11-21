const NickRegistery = require('./nick-registery');
const { buildMessage } = require('./utils');

const ValidNameRegex = /^[\w\d][\w\d_-]*$/i;

function updateRoomMembership(user, getRoom, payload, callback) {
    const serverRoom = getRoom();
    const targetRoom = getRoom(payload.msg);
    if (!targetRoom.roomName.match(ValidNameRegex) || targetRoom === serverRoom) {
        return Promise.reject(new Error(`${targetRoom.roomName} is not a valid room name`));
    }

    return callback(targetRoom);
}

/**
 * Each method name below is camelized version of command
 * e.g. set-nick => setNick, join-group => joinGroup etc.
 */
module.exports = {
    setNick(user, getRoom, payload) {
        const oldNick = user.nick;
        if (!payload.msg.match(ValidNameRegex)) {
            return Promise.reject(new Error(`${payload.msg} is an invalid nick`));
        }

        // Set new nick
        const newNick = NickRegistery.update(user.id, payload.msg);
        if (!newNick) {
            return Promise.reject(new Error(`Unable to change nick to ${payload.msg}`));
        }

        // Let user know
        const resp = buildMessage('nick-set', {
            oldNick,
            newNick,
            utc_timestamp: 0
        });
        user.nick = newNick;
        user.send(resp);

        user.joinedRooms.map(r => getRoom(r.roomName)).forEach(r => {
            r.publish(buildMessage('member-nick-set', {
                from: oldNick,
                to: r.roomName,
                pack_msg: resp.parsed
            }));
        });

        return Promise.resolve(resp);
    },

    joinGroup(user, getRoom, payload) {
        return updateRoomMembership(
            user,
            getRoom,
            payload,
            room => Promise.resolve(user.join(room))
        );
    },

    leaveGroup(user, getRoom, payload) {
        return updateRoomMembership(
            user,
            getRoom,
            payload,
            room => Promise.resolve(user.leave(room))
        );
    },

    listGroup(user, getRoom, payload) {
        const commandRoom = getRoom();
        const targetRoom = getRoom(payload.to);
        const nicks = targetRoom.userIds
            .map(id => NickRegistery.read(id))
            .filter(n => !!n);
        const msg = buildMessage('group-list', {
            to: targetRoom.roomName,
            from: commandRoom.roomName,
            pack_msg: nicks
        });

        user.send(msg);
        return Promise.resolve(msg);
    },

    sendMsg(user, getRoom, payload) {
        const commandRoom = getRoom();
        if (payload.to === commandRoom.roomName) {
            return Promise.reject(new Error(`${payload.to} doesn't support messages`));
        }

        const room = getRoom(payload.to);
        const resp = buildMessage('group-message', {
            to: payload.to,
            from: user.nick,
            msg: payload.msg
        });

        room.publish(resp);
        return Promise.resolve(resp);
    },

    ping(user) {
        user.send(buildMessage('pong'));
        user.lastSeen = new Date().getTime();
        return Promise.resolve();
    }
};
