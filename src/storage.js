const bluebird = require('bluebird');
const debug = require('debug')('raspchat:message-log');
const sqlite = require('sqlite3').verbose();
const $info = Symbol('info');

const MESSAGES_DDL = `
PRAGMA journal_mode=WAL;
CREATE TABLE IF NOT EXISTS chat_log (
    id      VARCHAR (255) PRIMARY KEY DESC
                          NOT NULL
                          UNIQUE,
    room_id VARCHAR (255) NOT NULL,
    message BLOB
);
CREATE INDEX IF NOT EXISTS chat_log_room_idx ON chat_log(room_id);
`;

const MESSAGES_UPSERT = `
INSERT OR REPLACE INTO chat_log(id, room_id, message)
VALUES(?, ?, ?);
`;

const MESSAGES_SELECT = `
SELECT * FROM chat_log
WHERE room_id = ? AND id >= ?
ORDER BY id DESC
LIMIT ?;
`;

const GIF_DDL = `
PRAGMA journal_mode=WAL;
CREATE TABLE IF NOT EXISTS gif (
    id      TEXT PRIMARY KEY
                 NOT NULL
                 UNIQUE,
    url     TEXT NOT NULL
);
`;

const GIF_UPSERT = `
INSERT OR REPLACE INTO gif(id, url)
VALUES (?, ?);
`;

const GIF_SELECT = `
SELECT * FROM gif
WHERE id = ?
`;

class MessageLog {
    constructor(db) {
        this[$info] = {
            db
        };
    }

    async open() {
        const db = this[$info].db;
        await db.execAsync(MESSAGES_DDL);
        debug('Chatlog database opened...');
    }

    async put(id, room_id, message) {
        debug('Saving message', id, room_id, message);
        const db = this[$info].db;
        await db.runAsync(MESSAGES_UPSERT, id, room_id, message);
    }

    async fetchFor(room_id, limit, start_id) {
        const db = this[$info].db;
        return await db.allAsync(MESSAGES_SELECT, room_id, start_id, limit);
    }
}

class GifStore {
    constructor(db) {
        this[$info] = { db };
    }

    async open() {
        const db = this[$info].db;
        await db.execAsync(GIF_DDL);
        debug('GIF database opened');
    }

    async put(keywords, url) {
        debug('Saving gif record');
        const db = this[$info].db;
        await db.runAsync(GIF_UPSERT, keywords, url);
    }

    async get(keywords) {
        const db = this[$info].db;
        return await db.allAsync(GIF_SELECT, keywords);
    }
}

const db = bluebird.promisifyAll(new sqlite.Database(process.env.DB_PATH || './chat-log.db'));
const ChatLog = new MessageLog(db);
const GifCache = new GifStore(db);

(async function () {
    await ChatLog.open();
    await GifCache.open();
})().catch(e => {
    console.error('Unable to open databases');
    console.error(e);
    process.exit(-1);
});

module.exports = {
    ChatLog,
    GifCache
};
