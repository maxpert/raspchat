import { DEFAULT_CHANNEL, MessageCommonProperties } from './message-commons';

const LogStateCommands = {
    'pong': m => null,
    'nick-set': m => null,

    'default': message => {
        console.error('Invalid message', message);
        return null;
    },

    'client-meta': m => ({
        ...MessageCommonProperties(m),
        to: DEFAULT_CHANNEL,
        from: DEFAULT_CHANNEL,
        type: 'meta',
        command: 'client-meta',
        message: m.msg
    }),

    [DEFAULT_CHANNEL.toLowerCase()]: m => ({ 
        ...MessageCommonProperties(m),
        to: DEFAULT_CHANNEL,
        from: DEFAULT_CHANNEL,
        message: m.msg,
        type: 'meta'
    }),

    'group-list': m => ({
        ...MessageCommonProperties(m),
        to: m.to,
        from: m.from,
        message: `members: ${m.pack_msg.join(', ')}`,
        type: 'meta'
    }),

    'group-join': m => ({ 
        ...MessageCommonProperties(m),
        to: m.to,
        from: m.from,
        message: `${m.from} has joined ${m.to}`,
        type: 'meta'
    }),

    'group-leave': m => ({ 
        ...MessageCommonProperties(m),
        to: m.to,
        from: m.from,
        message: `${m.from} has left ${m.to}`,
        type: 'meta'
    }),

    'member-nick-set': m => ({
        ...MessageCommonProperties(m),
        to: m.to,
        from: DEFAULT_CHANNEL,
        message: `${m.pack_msg.oldNick} changed nick to ${m.pack_msg.newNick}`,
        type: 'meta'
    }),

    'group-message': m => ({
        ...MessageCommonProperties(m),
        to: m.to,
        from: m.from,
        message: m.msg
    })
};

export default function(logs, message) {
    const command = message['@'].toLowerCase();
    const targetChannel = message.to || DEFAULT_CHANNEL;
    const channelLog = logs[targetChannel] || [];
    const parser = LogStateCommands[command] || LogStateCommands['default'];
    const parsedMessage = parser(message);

    if (!parsedMessage) {
        return { [targetChannel]: channelLog };
    }

    return {
        [targetChannel]: channelLog.concat([{
            ...parsedMessage,
            _raw: message
        }])
    };
}
