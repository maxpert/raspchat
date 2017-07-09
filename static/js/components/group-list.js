/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

var vue = require('vue');

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
    this.$on("message_new", this.newMessage);
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
