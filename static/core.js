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
        return fn.apply(scope);
    };
  };
}

window.core = (function(win, doc) {
  var MSG_DELIMETER = "~~~~>";
  var MSG_SENDER_REGEX = /^([^@]+)@+(.+)$/i;
  var SERVER_ALIAS = 'SERVER';

  var validCommandRegex = /^\/([a-zA-Z]+)\s*(.*)$/i;
  var nickRegex = /^\/nick\s+([A-Za-z0-9]+)$/i;
  var gifRegex = /^\/gif\s+(.+)$/i;
  var joinRegex = /^\/join\s+([A-Za-z0-9]+)$/i;
  var leaveRegex = /^\/leave\s+([A-Za-z0-9]+)$/i;
  var switchRegex = /^\/switch\s+([A-Za-z0-9]+)$/i;

  var giffer = {
    search: function (keywords, url_callback) {
      keywords = encodeURIComponent(keywords);
      var oReq = new XMLHttpRequest();
      oReq.addEventListener("load", function (r) {
        var resp = JSON.parse(oReq.response);
        console.log("gif info", resp);
        if (resp.attachments && resp.attachments.length > 0){
          var atch = resp.attachments[0];
          url_callback(atch.image_url, atch);
          return;
        }

        url_callback(null, null);
      });
      oReq.open("get", "/gif?q="+keywords);
      oReq.send();
    }
  };

  var parseRecipient = function (msg) {
    var matches = msg.match(MSG_SENDER_REGEX);
    if (matches && matches.length == 3){
      return {from: matches[1], to: matches[2]};
    }

    if (msg != SERVER_ALIAS) {
      return null;
    }

    return {from: SERVER_ALIAS, to: SERVER_ALIAS};
  };

  var parseMessage = function (data) {
    var msg = data.split(MSG_DELIMETER);
    if (msg.length < 2) {
      return null;
    }

    var p = parseRecipient(msg[0]);
    if (p == null) {
      return null;
    }

    p.msg = msg[1];
    return p;
  };

  var processComand = function (cmd, callback) {
    var match = cmd.match(validCommandRegex);

    if (!match){
      return false;
    }

    if (match[1].toLowerCase() == 'help') {
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
    fire: function (channel, data) {
      var subscribes = this._channels[channel] || [],
          l = subscribes.length;
      while (l--) {
        var cb = subscribes[l];
        win.setTimeout(function () {
          cb && cb.apply(this, data || [])
        }, 0);
      }
    },

    off: function (channel, handler) {
      var subscribes = this._channels[channel] || [],
          l = subscribes.length;

      while (l--) {
        if (subscribes[l] === handle) {
          subscribes.splice(l, 1);
        }
      }
    },

    on: function (channel, handler) {
      (this._channels[channel] = this._channels[channel] || []).push(handler);
    }
  };

  var Transport = function (url) {
    this.events = new EventEmitter();
    this.sock = null;
    this.id = null;
    this.url = url || 'http:///';
  };

  Transport.prototype = {
      connect: function () {
        this.sock = io(this.url);
        this.sock.on('connect', this._on_connect.bind(this));
        this.sock.on('disconnect', this._on_disconnect.bind(this));
        this.sock.on('new-msg', this._on_message.bind(this));
        this.sock.on('group-message', this._on_message.bind(this));
        this.sock.on('group-join', this._on_group_joined.bind(this));
        this.sock.on('group-leave', this._on_group_left.bind(this));
        this.sock.on('nick-set', this._on_nick_changed.bind(this));
        this.events.fire('connecting');
      },

      setNick: function (nick) {
        this.send(SERVER_ALIAS, "/nick "+nick);
      },

      send: function (to, msg) {
        var me = this;
        var processed = processComand(msg, function(cmd, value){
          if (cmd == "switch-group") {
            me.events.fire('switch', [value]);
            return;
          }

          var m = value;
          if (cmd == "send-msg") {
            m = to + MSG_DELIMETER + m;
          }
          me.sock.emit(cmd, m);
        });

        if (!processed) {
          this.sock.emit("send-msg", to + MSG_DELIMETER + msg);
        }
      },

      _on_message: function (msg) {
        var m = parseMessage(msg);
        if (m == null) {
          console.warn("Invalid data", msg);
          return;
        }

        m.delivery_time = new Date();
        this.events.fire('message', [m]);
      },

      _on_group_joined: function (msg) {
        var recpInfo = parseRecipient(msg);
        this.events.fire('message', [{
          from: SERVER_ALIAS,
          to: SERVER_ALIAS,
          delivery_time: new Date(),
          msg: recpInfo.from + " joined " + recpInfo.to
        }]);
        this.events.fire('joined', [recpInfo]);
      },

      _on_group_left: function (msg) {
        var recpInfo = parseRecipient(msg);
        this.events.fire('message', [{
          from: SERVER_ALIAS,
          to: SERVER_ALIAS,
          delivery_time: new Date(),
          msg: recpInfo.from + " left " + recpInfo.to
        }]);
        this.events.fire('leave', [recpInfo]);
      },

      _on_nick_changed: function (nick) {
        this.id = nick;
        this.events.fire('nick-changed', [nick]);
      },

      _on_connect: function () {
        this.id = this.sock.id;
        this.events.fire('connected');
      },

      _on_disconnect: function () {
        this.events.fire('disconnected');
      },
  };

  return {
    Transport: Transport,
    EventEmitter: EventEmitter,
  };
})(window, window.document);
