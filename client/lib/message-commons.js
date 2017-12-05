export const DEFAULT_CHANNEL = 'SERVER';

export const MessageCommonProperties = m => ({
    id: m['!id'],
    timestamp: m.utc_timestamp,
    command: m['@'].toLowerCase(),
});
