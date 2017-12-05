import messageParser from '../lib/message-parser';

export default {
    add: (state, actions) => message => messageParser(state, message),
    clear: (state, actions) => name => ({
        ...state,
        [name]: null
    })
};
