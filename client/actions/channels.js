const DEFAULT_CHANNEL = 'SERVER';
const ValidCommands = new Set(
    'nick-set', 
    'group-join', 
    'group-leave',
    'pong',
    'group-message'
);

export default {
    connect: (state, actions) => config => {
        const connection = new WebSocket(config.webSocketConnectionUri);
        connection.onopen = actions.onOpen(connection);
        connection.onclose = actions.onClose(connection);
        connection.onmessage = actions.onMessage(connection);
    },

    onOpen: (state, actions) => connection => event => {
        const connected = true;
        return {connection, connected};
    },

    onClose: (state, actions) => connection => event => {
        const connected = false;
        return {connection: null, connected};
    },

    onMessage: (state, actions) => connection => event => {
        try {
            const msg = JSON.parse(event.data);
            return actions.addMessage(msg);
        } catch (e) {
            console.warn('Error while parsing message', event.data);
        }

        return {};
    },

    addMessage: (state, actions) => message => {
        const command = message['@'].toLowerCase();
        const targetChannel = message.to || DEFAULT_CHANNEL;
        const channelLog = state.logs[targetChannel] || [];
        channelLog.push(message);
        return { 
            logs: { 
                [targetChannel]: channelLog 
            } 
        };
    }
};
