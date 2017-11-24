export default {
    toggleSound: (state, actions) => value => {
        const soundEnabled = value === undefined ? !state.soundEnabled : value;
        return Promise.resolve({ soundEnabled }).then(actions.applyState);
    },

    toggleNotification: (state, actions) => value => {
        const notificationEnabled = value === undefined ? !state.soundEnabled : value;
        return Promise.resolve({ notificationEnabled }).then(actions.applyState);
    },

    applyState: (state, actions) => state => state
};
