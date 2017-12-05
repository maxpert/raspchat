import view from './chat.view';
import logs from './chat.logs';
import * as builder from '../lib/message-builder';
import UserInputParser from '../lib/user-input-parser';

const HelpMessage = `
\`\`\`
System connected, valid commands:

/join <room> - to join a room
/leave <room> - to leave a room 
/nick <name> -  to set your nick
/list - to list room members

\`\`\`

# \u00A0
# Welcome to RaspChat
\u00A0\u00A0\u00A0
`;

export default {
    view,
    logs,

    connect: (state, actions) => config => {
        const connection = new WebSocket(config.webSocketConnectionUri);
        connection.onopen = actions.onOpen(connection);
        connection.onclose = actions.onClose(connection);
        connection.onmessage = actions.onMessage(connection);

        return { connectionAttempts: state.connectionAttempts + 1, connection };
    },

    sendMessageObject: (state, actions) => message => {
        if (!state.connection) {
            return {};
        }

        try {
            state.connection.send(JSON.stringify(message));
        } catch (e) {
            console.error('Unable to send message', e);
        }

        return {};
    },

    sendToRoom: (state, actions) => ({name: roomName, text}) => {
        return Promise.resolve(UserInputParser(roomName, text))
            .then(m => {
                actions.view.reset(roomName);
                actions.sendMessageObject(m);
                return m;
            })
            .then(({'@': command, 'msg': group}) => {
                if (command === 'leave-group') {
                    actions.logs.clear(group);
                    actions.view.clear(group);
                }
            });
    },

    stayConnected: (state, actions) => ({timeout, config}) => () => {
        const c = state.connection;
        const alreadyConnected = c && (c.readyState === 1 ||c.readyState === 0);

        // Do not retry if already connected or a connection was never attempted
        if (state.connectionAttempts === 0 || alreadyConnected) {
            window.setTimeout(actions.stayConnected({config, timeout}), timeout);
            return {};
        }

        const ret =  actions.connect(config);
        window.setTimeout(actions.stayConnected({config, timeout}), timeout);
        return ret;
    },

    onOpen: (state, actions) => connection => event => {
        actions.logs.add(builder.meta(HelpMessage));
        window.setTimeout(() => {
            actions.sendMessageObject(builder.join('Lounge'));
        }, 1);
        return { connectionAttempts: 1, connection };
    },

    onClose: (state, actions) => connection => event => {
        console.error('Connection closed', event);
        return { connection };
    },

    onMessage: (state, actions) => connection => event => {
        try {
            const msg = JSON.parse(event.data);
            actions.logs.add(msg);
        } catch (e) {
            console.warn('Error while parsing message', e, event.data);
        }

        return { connection };
    }
};
