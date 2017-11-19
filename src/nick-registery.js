const NameGen = require('project-name-generator');
const {genId} = require('./utils');

const NickIdMap = new Map();

const NickRegistry = {
    create: function (id) {
        let nick = NameGen({number: false}).raw.join('-');
        while (NickIdMap.has(`@${nick}`)) {
            nick = NameGen({number: true}).raw.join('-');
        }

        NickIdMap.set(`@${nick}`, id);
        NickIdMap.set(id, nick);
        return nick;
    },

    read: function (id) {
        if (NickIdMap.has(id)) {
            return NickIdMap.get(id);
        }

        return null;
    },

    delete: function (id) {
        if (!NickIdMap.has(id)) {
            return false;
        }

        const nick = NickIdMap.get(id);
        if (NickIdMap.has(`@${nick}`)) {
            NickIdMap.delete(`@${nick}`);
        }

        NickIdMap.delete(id);
        return true;
    },

    update: function (id, nick) {
        // Nick is already taken
        if (NickIdMap.has(`@${nick}`) && NickIdMap.get(`@${nick}`) !== id) {
            nick = nick + '-' + genId({short: true});
        }

        if (!NickRegistry.delete(id)) {
            return null;
        }

        NickIdMap.set(`@${nick}`, id);
        NickIdMap.set(id, nick);

        return nick;
    }
};

module.exports = NickRegistry;
