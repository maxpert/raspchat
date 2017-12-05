import * as builder from './message-builder';

const commandRegex = /^\/(join|leave|nick)\s+(.+)$/i;
const commandTransformMap = {
    'join': r => builder.join(r),
    'leave': r => builder.leave(r),
    'nick': r => builder.nick(r)
};

export default function (roomName, text) {
    const regexMatch = commandRegex.exec(text);
    if (regexMatch && commandTransformMap[regexMatch[1]]) {
        return commandTransformMap[regexMatch[1]](regexMatch[2].trim());
    }

    const trimmedText = text.trim().toLowerCase();
    if (trimmedText === '/leave') {
        return builder.leave(roomName);
    }

    if (trimmedText === '/list') {
        return builder.list(roomName);
    }

    // Parse rest of commands here...
    return builder.message(roomName, text);   
}
