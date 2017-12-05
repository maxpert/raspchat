import { DEFAULT_CHANNEL } from './message-commons';

let messageId = 0;

function serverCommandMessage(command, msg) {
    return {
        '@': command,
        to: DEFAULT_CHANNEL,
        msg
    };
}

export function meta(msg, extra) {
    extra = extra || {};

    return {
        '@': 'client-meta',
        '!id': `m${messageId++}`,
        'utc_timestamp': new Date().getTime(),
        to: DEFAULT_CHANNEL,
        msg,
        ...extra
    };
}

export function join(room, extra) {
    extra = extra || {};
    return {
        ...serverCommandMessage('join-group', room),
        ...extra
    };
}

export function leave(room, extra) {
    extra = extra || {};
    return {
        ...serverCommandMessage('leave-group', room),
        ...extra
    };
}

export function nick(nick, extra) {
    extra = extra || {};
    return {
        ...serverCommandMessage('set-nick', nick),
        ...extra
    };
}

export function list(room, extra) {
    extra = extra || {};
    return {
        ...serverCommandMessage('list-group', room),
        ...extra
    };
}

export function message(room, message, extra) {
    extra = extra || {};
    return {
        '@': 'send-msg',
        to: room,
        msg: message,
        ...extra
    };
}
