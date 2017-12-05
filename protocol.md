
# Protocol

Every message is just a JSON object; Each pay load must have 2 fields `@` which is command name; `!id` which is a unique message id. In the examples described all placeholders are described with enclosing angle brackets `<<...>>`. Following placeholders are fixed:

 - `<<id>>`: `{String}` is unique message ID; is always supposed to be incrementing id
 - `<<ts>>`: `{Number}` unix timestamp
 - `<<nick>>`: `{String}` Nick name of user
 - `<<group_id>>`: `{String}` Name of group
 - `<<payload>>`: `{String}` message payload from server 

### SERVER meta command

Everything server wants to send is wrapped under SERVER meta command. This might include status messages, meta payloads (channel user list, banned users etc.)

```json
{"@":"SERVER","!id":<<id>>,"utc_timestamp":0,"msg":"<<payload>>"}
```

### Nick set

 - `<<nick>>`: `{String}` Nick name user is trying to set

```json
{"@":"set-nick","to":"SERVER","msg":"<<nick>>"}
{"@":"nick-set","!id":<<id>>,"utc_timestamp":0,"oldNick":"<<nick>>","newNick":"<<nick>>"}
{"@":"member-nick-set","!id":<<id>>,"utc_timestamp":<<ts>>,"from":<<nick>>,"to":<<group_id>>,"pack_msg":{"@":"nick-set","!id":<<id>>,"utc_timestamp":0,"oldNick":<<nick>>,"newNick":<<nick>>}}
```

### Join/Leave group

```json
{"@":"join-group","to":"SERVER","msg":"<<group_id>>"}
{"@":"leave-group","to":"SERVER","msg":"<<group_id>>"}
{"@":"group-join","!id":<<id>>,"utc_timestamp":<<ts>>,"to":"<<group_id>>","from":"<<nick>>"}
{"@":"group-leave","!id":<<id>>,"utc_timestamp":<<ts>>,"to":"<<group_id>>","from":"<<nick>>"}
```

### Ping/Pong

```json
{"@":"ping","!id":<<id>>,"utc_timestamp":0,"t":<<ts>>}
{"@":"pong","t":<<ts>>}
```

### Send/Receive message

```json
{"@":"send-msg","to":"<<group_id|'@'+nick>>","msg":"<<payload>>"}
{"@":"group-message","!id":<<id>>,"utc_timestamp":<<ts>>,"to":"<<group_id|'@'+nick>>","from":"<<nick>>","msg":"<<payload>>"}
```
