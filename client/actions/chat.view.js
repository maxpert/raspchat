import messageParser from '../lib/message-parser';

const DefaultViewState = { active: false, userInput: '' };

export default {
    init: (state, actions) => name => ({
        ...state,
        [name]: { ... DefaultViewState }
    }),

    clear: (state, actions) => name => ({
        ...state,
        [name]: null
    }),

    reset: (state, actions) => name => actions.update({name, userInput: ''}),

    update: (state, actions) => ({name, userInput, active}) => {
        const old = (state[name] || {...DefaultViewState});
        if (userInput === undefined) userInput = old.userInput;
        if (active === undefined) active = old.active;

        const newState = {
            ...state,
            [name]: { userInput, active }
        };
        
        return newState;
    }
};
