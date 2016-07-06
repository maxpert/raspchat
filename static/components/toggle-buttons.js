/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

(function (vue, win, doc) {
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
})(Vue, window, window.document);
