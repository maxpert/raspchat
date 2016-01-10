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

window.core = (function(win, doc) {
  var SERVER_ALIAS = 'SERVER';

  var validCommandRegex = /^\/([a-zA-Z]+)\s*(.*)$/i;
  var nickRegex = /^\/nick\s+([A-Za-z0-9]+)$/i;
  var gifRegex = /^\/gif\s+(.+)$/i;
  var joinRegex = /^\/join\s+([A-Za-z0-9]+)$/i;
  var leaveRegex = /^\/leave\s+([A-Za-z0-9]+)$/i;
  var listRegex = /^\/list\s+([A-Za-z0-9]+)$/i;
  var switchRegex = /^\/switch\s+([A-Za-z0-9]+)$/i;

  var giffer = {
    search: function (keywords, url_callback) {
      keywords = encodeURIComponent(keywords);
      var oReq = new XMLHttpRequest();
      oReq.addEventListener("load", function (r) {
        var resp = JSON.parse(oReq.response);
        if (resp && resp.url){
          url_callback(resp.url, resp);
          return;
        }

        url_callback(null, null);
      });
      oReq.open("get", "/gif?q="+keywords);
      oReq.send();
    }
  };

  var processComand = function (cmd, callback) {
    var match = cmd.match(validCommandRegex);

    if (!match){
      return false;
    }

    if (match[1].toLowerCase() == 'help') {
      return true;
    }

    match = cmd.match(listRegex);
    if (match) {
      callback("list-group", (match[1] && match[1].trim()) || SERVER_ALIAS);
      return true;
    }

    match = cmd.match(nickRegex);
    if (match) {
      callback("set-nick", match[1]);
      return true;
    }

    match = cmd.match(switchRegex);
    if (match) {
      callback("switch-group", match[1]);
      return true;
    }

    match = cmd.match(joinRegex);
    if (match) {
      callback("join-group", match[1]);
      return true;
    }

    match = cmd.match(leaveRegex);
    if (match) {
      callback("leave-group", match[1]);
      return true;
    }

    match = cmd.match(gifRegex);
    if (match) {
      giffer.search(match[1], function (url, obj) {
        var t = cmd;
        if (url) {
          t = "!["+cmd+"]("+url+")";
        }

        callback("send-msg", t);
      });

      return true;
    }

    return false;
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
    this.url = url || ('ws://'+win.location.host+'/chat');
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

      send: function (to, msg) {
        var me = this;
        var processed = processComand(msg, function(cmd, value){
          if (cmd == "switch-group") {
            me.events.fire('switch', value);
            return;
          }

          var m = {'@': cmd, to: to, msg: value};
          me.sock.send(JSON.stringify(m));
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
          console.error(e);
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
          console.error("Error decoding", e.data, er);
        }

        if (data['@']) {
          this._handleMessage(data);
        }
      },

      _completeHandShake: function (msg) {
        if (!this.handshakeCompleted) {
          this.handshakeCompleted = true;
          this.events.fire('handshake', SERVER_ALIAS);
          this.events.fire('message', {
            from: SERVER_ALIAS,
            to: SERVER_ALIAS,
            msg: "```"+msg.msg+"```",
          });
          this.setNick(this.nick);
        }
      },

      _handleMessage: function (msg) {
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
        this.events.fire('message', {
          from: SERVER_ALIAS,
          to: SERVER_ALIAS,
          delivery_time: new Date(),
          msg: msg.from + " joined " + msg.to
        });
        this.events.fire('joined', msg);
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

  var _globalTransport = {};

  return {
    Transport: Transport,
    EventEmitter: EventEmitter,
    GetTransport: function (name, url) {
      _globalTransport[name] = _globalTransport[name] || new Transport(url);
      return _globalTransport[name];
    },
  };
})(window, window.document);
