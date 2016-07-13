/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

if (!Function.prototype.bind){
  Function.prototype.bind = function (scope) {
    var fn = this;
    return function () {
        var args = Array.prototype.slice.call(arguments);
        return fn.apply(scope, args);
      };
  };
}

window.$glueFunctions = function (obj) {
  for (var i in obj) {
    if (obj[i] instanceof Function) {
      obj[i] = obj[i].bind(obj);
    }
  }
};

window.$mix = function (){
    var ret = {};
    if (Object.assign) {
      var args = Array.prototype.slice.call(arguments);
      return Object.assign.apply(Object, args);
    }

    for(var i=0; i<arguments.length; i++)
        for(var key in arguments[i])
            if(arguments[i].hasOwnProperty(key))
                ret[key] = arguments[i][key];
    return ret;
};

window.core = (function(win, doc, raspconfig) {
  var SERVER_ALIAS = 'SERVER';

  var validCommandRegex = /^\/(nick|gif|join|leave|list|switch)\s*(.*)$/i;
  var userCommandToEventsMap = {
    'list': {eventName: 'list-group', paramRequired: false, defaultParam: false},
    'nick': {eventName: 'set-nick', paramRequired: true},
    'switch': {eventName: 'switch-group', paramRequired: true},
    'join': {eventName: 'join-group', paramRequired: true},
    'leave': {eventName: 'leave-group', paramRequired: false, defaultParam: false},
    'gif': {eventName: 'send-gif', paramRequired: true}
  };
  var processComand = function (cmd, callback) {
    var match = cmd.match(validCommandRegex);

    // map should have command
    if (!match || !userCommandToEventsMap[match[1]]){
      return false;
    }

    if (match[1].toLowerCase() == 'help') {
      return true;
    }

    // Invoke matched command
    var selectedCmd = userCommandToEventsMap[match[1]];
    var cmdParam = match[2];
    if (selectedCmd.paramRequired && !cmdParam) {
      return false;
    }
    else
    {
      cmdParam = cmdParam || selectedCmd.defaultParam;
    }

    callback(selectedCmd.eventName, cmdParam);
    return true;
  };

  var getWebSocketConnectionUri = function () {
    var loc = win.location;
    var isSecure = loc.protocol.toLowerCase().endsWith("s:");
    var wsUri = raspconfig && raspconfig.webSocketConnectionUri;
    var wssUri = raspconfig && raspconfig.webSocketSecureConnectionUri;
    var templateString = (isSecure ? wssUri : wsUri) || "{protocol}//{host}/chat";
    var resultString = ""+templateString;
    var replacableHrefProperties = {
      "host": "{host}",
      "port": "{port}",
      "hostname": "{hostname}",
      "pathname": "{path}"
    };

    for (var property in replacableHrefProperties) {
      var placeHolder = replacableHrefProperties[property];
      if (resultString.indexOf(placeHolder) > -1) {
        resultString = resultString.replace(placeHolder, ""+loc[property]);
      }
    }

    resultString = resultString.replace("{protocol}", isSecure ? "wss:" : "ws:");
    return resultString;
  };

  var EventEmitter = function () {
    this._channels = {};
  };

  EventEmitter.prototype = {
    fire: function (channel) {
      var subscribes = this._channels[channel] || [],
          l = subscribes.length,
          data = Array.prototype.slice.call(arguments, 1);
      for (var i = 0; i < l; i++) {
        (function (j) {
          var cb = subscribes[j];
          win.setTimeout(function () {
            cb && cb.apply(this, data || [])
          }, 0);
        })(i);
      }
    },

    off: function (channel, handler) {
      var subscribes = this._channels[channel] || [],
          l = subscribes.length;

      while (l--) {
        if (subscribes[l] === handler) {
          subscribes.splice(l, 1);
        }
      }
    },

    on: function (channel, handler) {
      this._channels[channel] = this._channels[channel] || [];
      this._channels[channel].push(handler);
    }
  };

  var Transport = function (url) {
    $glueFunctions(this);
    this.events = new EventEmitter();
    this.sock = null;
    this.handshakeCompleted = false;
    this.url = url || getWebSocketConnectionUri();
  };

  Transport.prototype = {
      connect: function (nick) {
        this.nick = nick;
        this._create_ws_connection();
      },

      setNick: function (nick) {
        this.send(SERVER_ALIAS, "/nick "+nick);
      },

      sendRaw: function (to, msg) {
        this.sock.send(JSON.stringify({"@": "send-raw-msg", to: to, msg: msg}));
      },

      isValidCmd: function (msg) {
        var match = msg.match(validCommandRegex);
        if (!match) {
          return false;
        }

        return true;
      },

      getHistory: function (grp, offset, limit) {
        offset = offset || 0;
        limit = limit || 50;
        var me = this;
        var encodedGroupName = encodeURIComponent(grp);
        atomic.setContentType('application/json');
        atomic.get("/chat/api/channel/"+encodedGroupName+"/message?limit="+limit+"&offset="+offset)
             .success(function (response, xhr) {
               me._on_group_history_recvd(grp, response);
             })
             .error(function (response, xhr) {
               events.fire('history-error', xhr);
             });
      },

      send: function (to, msg) {
        var me = this;
        var processed = processComand(msg, function(cmd, cmdParam){
          if (cmd == "switch-group") {
            me.events.fire('switch', cmdParam);
            return;
          }

          if (cmd == "send-gif") {
            giffer.search(cmdParam, function (url, obj) {
              var t = msg;
              if (url) {
                t = "> !["+cmdParam+"]("+url+")\n\n> **GIF** "+cmdParam+"\n";
              }

              me.sock.send(JSON.stringify({"@": "send-msg", to: to, msg: t}));
            });

            return;
          }

          // Populate /leave <group-name> if <group-name> was not provided
          // Populate /list <group-name> if <group-name> was not provided
          if (cmd == "leave-group" || cmd == "list-group") {
            cmdParam = cmdParam || to;
          }

          me.sock.send(JSON.stringify({'@': cmd, to: to, msg: cmdParam}));
        });

        if (!processed) {
          this.sock.send(JSON.stringify({"@": "send-msg", to: to, msg: msg}));
        }
      },

      _create_ws_connection: function () {
        try{
          if (this.sock && this.sock.close) {
            this.sock.onclose = null;
            this.sock.onopen = null;
            this.sock.onmessage = null;
            this.sock.close();
          }
        }catch(e){
          console && console.error(e);
        }

        this.sock = new WebSocket(this.url);
        this.sock.onopen = this._on_connect;
        this.sock.onclose = this._on_disconnect;
        this.sock.onmessage = this._on_data;
        this.events.fire('connecting');
      },

      _on_connect: function (e) {
        this.events.fire('connected');
      },

      _on_disconnect: function () {
        this.handshakeCompleted = false;
        this.events.fire('disconnected');
        var me = this;
        win.setTimeout(function (){ me._create_ws_connection(); }, 1000);
      },

      _on_data: function (e) {
        var data = {};
        try {
          data = JSON.parse(e.data);
        }catch(er){
          console && console.error("Error decoding", e.data, er);
        }

        if (data['@']) {
          this._handleMessage(data);
        }
      },

      _completeHandShake: function (msg) {
        if (!this.handshakeCompleted) {
          this.handshakeCompleted = true;
          this.setNick(this.nick);
          this.events.fire('handshake', SERVER_ALIAS);
          this.events.fire('message', {
            from: SERVER_ALIAS,
            to: SERVER_ALIAS,
            msg: "```"+msg.msg+"```",
          });
        }
      },

      _handleMessage: function (msg) {

        // Switch case for handling message types
        // Ideal is to create a map and invoke methods directly
        switch (msg['@']) {
          case SERVER_ALIAS:
            this._completeHandShake(msg);
            break;

          case 'group-join':
            this._on_group_joined(msg);
            break;

          case 'group-leave':
            this._on_group_left(msg);
            break;

          case 'group-message':
            this._on_message(msg);
            break;

          case 'nick-set':
            this._on_nick_changed(msg);
            break;

          case 'member-nick-set':
            this._on_member_nick_changed(msg.to, msg.pack_msg);
            break;

          case 'group-list':
            this._on_group_members_list(msg.to, msg.pack_msg);
            break;

          case 'new-raw-msg':
            this._on_rawmessage(msg.to, msg.pack_msg);
            break;

          case 'ping':
            this.sock.send(JSON.stringify({'@': 'pong', t: msg.t}));
            break;

          default:
            break;
        }
      },

      _on_rawmessage: function (from, msg) {
        this.events.fire('raw-message', from, msg);
      },

      _on_message: function (msg) {
        msg.delivery_time = msg.delivery_time || new Date();
        this.events.fire('message', msg);
      },

      _on_group_joined: function (msg) {
        var events = this.events;
        events.fire('message', {
          from: SERVER_ALIAS,
          to: SERVER_ALIAS,
          delivery_time: new Date(),
          msg: msg.from + " joined " + msg.to
        });

        events.fire('joined', msg);
        this.getHistory(msg.to);
      },

      _on_group_history_recvd: function (grp, hist) {
        var historyMessages = hist.messages.map(this._prepareMetaMessage).reverse();
        this.events.fire('history', $mix(hist, {messages: historyMessages}));
      },

      _prepareMetaMessage: function (msg) {
        var ret = $mix(msg);
        switch (msg['@']) {
          case 'group-join':
            ret.meta = {action: 'joined'};
            break;
          case 'group-leave':
            ret.meta = {action: 'leave'};
            break;
          default:
            ret.meta = null;
        }

        return ret;
      },

      _on_group_members_list: function (to, list) {
        this.events.fire('members-list', to, list);
      },

      _on_group_left: function (recpInfo) {
        this.events.fire('message', {
          from: SERVER_ALIAS,
          to: SERVER_ALIAS,
          delivery_time: new Date(),
          msg: recpInfo.from + " left " + recpInfo.to
        });
        this.events.fire('leave', recpInfo);
      },

      _on_nick_changed: function (msg) {
        this.nick = msg.newNick;
        this.events.fire('nick-changed', msg.newNick, msg.oldNick);
      },

      _on_member_nick_changed: function (group, nickInfo) {
        this.events.fire('message', {
          to: group,
          msg: nickInfo.oldNick + " changed nick to " + nickInfo.newNick,
          from: nickInfo.newNick,
          delivery_time: new Date(),
        });
      },
  };

  Transport.HelpMessage = "Valid commands are: \n"+
            "/help for this help :)\n" +
            "/list for list of members in a group\n"+
            "/gif <gif-keywords> to send a gif \n"+
            "/join <group_name> to join a group (case-sensitive)\n"+
            "/nick <new_name> for changing your nick (case-sensitive)\n"+
            "/switch <group_name> to switch to a joined group (case-sensitive)\n";

  var _globalTransport = {};

  var giffer = {
    search: function (keywords, url_callback) {
      keywords = encodeURIComponent(keywords);
      atomic.setContentType('application/json');
      atomic.get("/gif?q="+keywords)
           .success(function (response, xhr) {
             url_callback(response.url, response);
           })
           .error(function (response, xhr) {
             url_callback(null, null);
           });
    }
  };

  return {
    Transport: Transport,
    EventEmitter: EventEmitter,
    GetTransport: function (name, url) {
      _globalTransport[name] = _globalTransport[name] || new Transport(url);
      return _globalTransport[name];
    },
  };
})(window, window.document, window.RaspConfig);
