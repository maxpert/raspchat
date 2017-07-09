/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

var vue = require('vue')
var win = window;
var doc = document;

var ToggleButtonMixin = {
  data: function () {
    return {enabled: false};
  },
  methods: {
    toggle: function () {
      var oldValue = this.$get("enabled");
      this.$set("enabled", !oldValue);
      this.onEnableChanged && this.onEnableChanged(oldValue); // jshint ignore: line 
    }
  }
};

vue.component('toast-notification-button', vue.extend({
  template: '#toast-notification-button',
  mixins: [ToggleButtonMixin],
  props: ["ignoreFor"],
  ready: function () {
    this.$set("enabled", Notification.permission === 'granted');
    this.$on("message_new", this.onNotification);
  },

  methods: {
    onNotification: function (msg, metaInfo) {
      if (!this.$get('enabled') ||
          metaInfo.noNotification ||
          msg.from == this.ignoreFor ||
          doc.hasFocus()) {
        return;
      }

      var bodyText =  ""+msg.msg;
      if (bodyText.length > 64) {
        bodyText = bodyText.substring(0, 64) + "...";
      }

      var notif = new Notification(msg.from, {body: bodyText, icon: "/static/favicon/favicon.ico"});
      notif.onclick = function () {
          win.focus();
          this.close();
      };
    },

    onEnableChanged: function (oldValue) {
      if (oldValue === false && Notification.permission !== 'granted') {
        Notification.requestPermission(this.onPermissionChanged);
      }
    },

    onPermissionChanged: function (permission) {
      if (permission !== 'granted' && this.$get('enabled')) {
        this.$set('enabled', false);
      }
    }
  }
}));

vue.component('sound-notification-button', vue.extend({
  template: '#sound-notification-button',
  mixins: [ToggleButtonMixin],
  props: ["defaultEnabled", "ignoreFor"],
  ready: function () {
    this.pingSound = new Audio("/static/ping.mp3");
    this.playingSound = false;

    if (this.defaultEnabled){
      this.$set("enabled", true);
    }

    this.$on("message_new", this.onNotification);
  },
  methods: {
    onNotification: function (msg, metaInfo) {
      if (this.playingSound) {
        return;
      }

      if (this.enabled && !metaInfo.noNotification && msg.from != this.ignoreFor){
        this.pingSound.play();
        this.playingSound = true;
        win.setTimeout(this.markPlayed, 1000);
      }
    },

    markPlayed: function () {
      this.playingSound = false;
    }
  }
}));
