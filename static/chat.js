/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

(function (vue) {
  var md = new markdownit("default", {
    linkify: true,
  });

  var showNotification = function (message) {
    if (!("Notification" in window)) {
      return false;
    } else if (Notification.permission === "granted") {
      var notification = new Notification(message);
    } else if (Notification.permission !== 'denied') {
      Notification.requestPermission(function (permission) {
        if (permission === "granted") {
          var notification = new Notification(message);
        }
      });
    }

    return true;
  };

  vue.filter('markdown', function (value) {
    return md.render(value);
  });

  vue.filter('better_date', function (value) {
    return moment(value).calendar();
  });

  vue.filter('avatar_url', function (value) {
    // http://api.adorable.io/avatars/face/eyes6/nose7/face1/AA0000
    return 'http://api.adorable.io/avatars/256/zmg-' + value + '.png';
  });

  vue.component('chat-message', vue.extend({
    props: ['message'],
    template: '#chat-message',
    ready: function () {
      this.$dispatch("chat-message-added", this.message);
      this.$watch("message.msg", function () {
        this.$dispatch("chat-message-added", this.message);
      }.$glue(this));
    }
  }));

  vue.component('chat-log', vue.extend({
    props: ['messages'],
    template: '#chat-messages',
    ready: function () {
      this.$el.addEventListener("click", function (event) {
        event = event || window.event;

        if (event.target.tagName == "A") {
          window.open(event.target.href, "_blank");
          event.preventDefault();
          event.stopPropagation();
        }
      }, false);
    },
    methods: {
      scrollToBottom: function () {
        this.$el.scrollTop = this.$el.scrollHeight;
      },
    },
  }));

  vue.component('chat-compose', vue.extend({
    template: '#chat-compose',
    data: function () {
      return {
        message: '',
      };
    },
    methods: {
      enterPressed: function (e) {
        var msg = this.message;
        if (e.shiftKey){
          this.$set('message', msg+'\n');
          return;
        }

        this.$set('message', '');
        this.$dispatch('send-message', msg);
      },

      tabPressed: function () {
        var msg = this.$get('message');
        this.$set('message', msg+'  ');
      },
    },
  }));

  vue.component('app-bar', vue.extend({
    props: ['userId'],
    template: '#app-bar',
    data: function () {
    },
    methods: {
    }
  }));

  vue.component('groups-list', vue.extend({
    template: '#groups-list',
    data: function () {
      return {
        groups: [],
        selected: "",
      };
    },
    ready: function () {
      this.groupsInfo = {};
      this.$on("group_joined", this.groupJoined);
      this.$on("group_switched", this.groupSwitch);
      this.$on("group_left", this.groupLeft);
      this.$on("message_new", this.newMessage)
    },
    methods: {
      selectGroup: function (id) {
        this._setUnread(id, 0);
        this.$set("selected", id);
        this.$dispatch("switch", id);
      },

      leaveGroup: function (id) {
        this.$dispatch("leave", id);
      },

      groupSwitch: function (group) {
        this.selectGroup(group);
      },

      groupJoined: function (group) {
        var groupInfo = this.groupsInfo[group] = this.groupsInfo[group] || {name: group, unread: 0, index: this.groups.length};
        this.groups.push(groupInfo);
      },

      groupLeft: function (group) {
        var g = this.groupsInfo[group] || {index: -1};
        if (g.index != -1){
          this.groups.splice(g.index, 1);
        }
      },

      newMessage: function (msg) {
        if (this.selected == msg.to || !this.groupsInfo[msg.to]) {
          return true;
        }

        this._setUnread(msg.to, this._getUnread(msg.to) + 1);
        return true;
      },

      _getUnread: function (g) {
        return (this.groupsInfo[g] && this.groupsInfo[g].unread) || 0;
      },

      _setUnread: function (g, count) {
        vue.set(this.groupsInfo[g], "unread", count);
        return true;
      }
    }
  }));

  var ToggleButtonMixin = {
    data: function () {
      return {enabled: false};
    },
    methods: {
      toggle: function () {
        this.$set("enabled", !this.$get("enabled"));
      }
    }
  };

  vue.component('sound-notification-button', vue.extend({
    template: '#sound-notification-button',
    mixins: [ToggleButtonMixin],
    props: ["defaultEnabled", "ignoreFor"],
    ready: function () {
      if (this.defaultEnabled){
        this.$set("enabled", true);
      }

      this.$on("message_new", this.onNotification);
    },
    methods: {
      onNotification: function (msg) {
        if (this.enabled && msg.from != this.ignoreFor){
          var snd = new Audio("/static/ping.mp3");
          snd.play();
        }
      }
    }
  }));

  var groupsLog = {};
  var vueApp = new vue({
    el: '#root',
    data: {
      nick: "",
      currentGroup: {name: '', messages: []},
      isConnected: false,
      isReady: false,
    },

    ready: function (argument) {
      this.transport = new core.Transport();
      this.transport.events.on('connected', this.onConnected);
      this.transport.events.on('disconnected', this.onDisconnected);
      this.transport.events.on('handshake', this.onHandshaked);

      this.transport.events.on('message', this.onMessage);
      this.transport.events.on('joined', this.onJoin);
      this.transport.events.on('leave', this.onLeave);
      this.transport.events.on('switch', this.onSwitch);
      this.transport.events.on('nick-changed', this.changeNick);

      this.$on("switch", this.onSwitch);
      this.$on("leave", function (group) {
        this.transport.send(group, "/leave "+group);
      });
      this.$watch("currentGroup.name", function (newVal, oldVal) {
        this.$broadcast("group_switched", newVal);
      });
    },

    methods: {
      connect: function () {
        this.$set("isConnecting", true);
        this.transport.connect(this.nick);
      },

      sendMessage: function (msg) {
        // Don't let user send message on default group
        if (this.currentGroup.name == this.defaultGroup && msg[0] != "/"){
          this._appendMetaMessage(
            this.defaultGroup,
            "You can only send a command here ...\n"+
            "Valid commands are: \n"+
            "/join <group_name> to join a group (case-sensitive)\n"+
            "/nick <new_name> for changing your nick (case-sensitive)\n"+
            "/switch <group_name> to switch to a joined group (case-sensitive)\n"
          );
          return;
        }

        this.transport.send(this.currentGroup.name, msg);
      },

      switchGroup: function (grp) {
        this.onSwitch(grp);
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
        if (!this._getGroupLog(group)) {
          alert('You have not joined group '+group);
          return true;
        }

        if (this.currentGroup.name == group) {
          return true;
        }

        this.$set('currentGroup.name', group);
        this.$set('currentGroup.messages', groupsLog[group]);
        return false;
      },

      _appendMessage: function (m) {
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

        this.$broadcast('message_new', m);
      },

      _appendMetaMessage: function (group, msg) {
        var groupLog = this._getOrCreateGroupLog(group);

        if (!this.currentGroup.name) {
          this.$set('currentGroup.name', group);
          this.$set('currentGroup.messages', groupLog);
        }

        groupLog.push({isMeta: true, msg: msg});
      },

      _getOrCreateGroupLog: function (g) {
        if (!groupsLog[g]) {
          groupsLog[g] = [];
          this.$broadcast("group_joined", g);
        }

        return groupsLog[g];
      },

      _getGroupLog: function (g) {
        return groupsLog[g] || null;
      }
    },
  });
})(Vue);
