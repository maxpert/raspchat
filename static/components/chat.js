/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

(function (vue, win, doc) {
  var groupsLog = {};
  var vueApp = new vue({
    el: '#root',
    data: {
      nick: "",
      currentGroup: {name: '', messages: []},
      isConnected: false,
      isConnecting: false,
      isReady: false,
      showAppBar: false,
    },

    ready: function () {
      if (this.$el.offsetWidth > 600){
        this.$set("showAppBar", true);
      }

      this.transport = core.GetTransport("chat");
      this.transport.events.on('connected', this.onConnected);
      this.transport.events.on('disconnected', this.onDisconnected);
      this.transport.events.on('handshake', this.onHandshaked);

      this.transport.events.on('raw-message', this.onRawMessage);
      this.transport.events.on('message', this.onMessage);
      this.transport.events.on('joined', this.onJoin);
      this.transport.events.on('leave', this.onLeave);
      this.transport.events.on('switch', this.onSwitch);
      this.transport.events.on('history', this.onHistoryRecv);
      this.transport.events.on('nick-changed', this.changeNick);
      this.transport.events.on('members-list', this.onMembersList);

      this.$on("switch", this.onSwitch);
      this.$on("leave", function (group) {
        this.transport.send(group, "/leave "+group);
      });

      this.$on("hamburger-clicked", function (v) {
        this.$set("showAppBar", !this.showAppBar);
      });

      this.$watch("currentGroup.name", function (newVal, oldVal) {
        this.$broadcast("group_switched", newVal);
      });
    },

    methods: {
      connect: function () {
        this.$set("isConnecting", true);
        this.$set("isConnected", true);
        this.transport.connect(this.nick);
      },

      sendMessage: function (msg) {
        // Don't let user send message on default group
        if (msg[0] == '/' && (!this.transport.isValidCmd(msg) || msg.toLowerCase().startsWith("/help")))
        {
          this._appendMetaMessage(this.currentGroup.name, core.Transport.HelpMessage);
          return;
        }

        this.transport.send(this.currentGroup.name, msg);
      },

      onRawMessage: function (from, msg) {
        if (msg.type != "Negotiate") {
          return;
        }

        this._appendMetaMessage(this.currentGroup.name, "DCC to "+from);
        var p = new core.PeerConnectionNegotiator(this.transport);
        p.events.on("close", function () {
          p.close();
        });
        p.connectTo(from);
      },

      switchGroup: function (grp) {
        this.onSwitch(grp);
      },

      onMembersList: function (group, list) {
        this._appendMessage({
          to: group,
          from: this.defaultGroup,
          msg: "Channel members for **"+group+"**\n\n - " + list.join("\n - "),
          delivery_time: new Date()
        });
      },

      onHandshaked: function (info_channel) {
        this.defaultGroup = info_channel;
        this.transport.send(this.defaultGroup, "/join lounge");
      },

      onMessage: function (m) {
        this._appendMessage(m);
      },

      onConnected: function () {
        this.$set('isConnected', true);
        this.$broadcast("connection_on");
      },

      changeNick: function (newNick) {
        this.$set('nick', newNick);
      },

      onDisconnected: function () {
        this.$set("isConnecting", true);
        this.$broadcast("connection_off");
      },

      onJoin: function (joinInfo) {
        this._getOrCreateGroupLog(joinInfo.to);
        this._appendMetaMessage(joinInfo.to, joinInfo.from + " has joined");
        if (this.currentGroup.name == this.defaultGroup) {
          this.switchGroup(joinInfo.to);
        }

        if (this.isConnecting) {
          this.$set("isConnecting", false);
        }
      },

      onLeave: function (info) {
        if (info.from == this.nick) {
          delete groupsLog[info.to];
          this.$broadcast("group_left", info.to);
        } else {
          this._appendMetaMessage(info.to, info.from + " has left");
        }

        if (this.currentGroup.name == info.to && this.nick == info.from) {
          this.switchGroup(this.defaultGroup);
        }
      },

      onSwitch: function (group) {
        if (this.$el.offsetWidth < 600) {
          this.$set("showAppBar", false);
        }

        if (!this._getGroupLog(group)) {
          alert('You have not joined group '+group);
          return true;
        }

        if (this.currentGroup.name == group) {
          return true;
        }

        this.$broadcast('group-switching', group);
        this.$set('currentGroup.name', group);
        this.$set('currentGroup.messages', groupsLog[group]);
        this.$broadcast('group-switched', group);
        return false;
      },

      onHistoryRecv: function (historyObj) {
        var msgs = historyObj.messages;

        this._clearGroupLogs();
        for (var i in msgs) {
          var m = msgs[i];
          if (!m.meta) {
            this._appendMessage(m, true);
          } else {
            switch (m.meta.action) {
              case 'joined':
                this.onJoin(m);
                break;
            }
          }
        }

        this.$broadcast('history-added', historyObj.id);
      },

      _appendMessage: function (m, silent) {
        var groupLog = this._getOrCreateGroupLog(m.to);

        if (!this.currentGroup.name) {
          this.$set('currentGroup.name', m.to);
          this.$set('currentGroup.messages', groupLog);
        }

        if (groupLog.length && groupLog[groupLog.length - 1].from == m.from) {
          var lastMsg = groupLog[groupLog.length - 1];
          lastMsg.msg += "\n\n" + m.msg;
        } else {
          groupLog.push(m);
        }

        this._limitGroupHistory(m.to);

        // no need
        if (silent) {
          return;
        }

        this.$broadcast('message_new', m, {noSound: m.to == this.defaultGroup});
      },

      _appendMetaMessage: function (group, msg) {
        var groupLog = this._getOrCreateGroupLog(group);

        if (!this.currentGroup.name) {
          this.$set('currentGroup.name', group);
          this.$set('currentGroup.messages', groupLog);
        }

        groupLog.push({isMeta: true, msg: msg});
        this._limitGroupHistory(group);
      },

      _limitGroupHistory: function (group, limit) {
        limit = limit || 100;
        var log = this._getOrCreateGroupLog(group);

        if (log.length > limit) {
          log.splice(0, log.length - limit);
        }
      },

      _getOrCreateGroupLog: function (g) {
        if (!groupsLog[g]) {
          groupsLog[g] = [];
          this.$broadcast("group_joined", g);
        }

        return groupsLog[g];
      },

      _clearGroupLogs: function (g) {
        var logs = this._getGroupLog(g);
        if (logs) logs.splice(0, logs.length);
      },

      _getGroupLog: function (g) {
        return groupsLog[g] || null;
      }
    },
  });
})(Vue, window, window.document);
