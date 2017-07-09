/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

var vue = require('vue');
var win = window;

vue.component('chat-log', vue.extend({
  props: ['messages'],
  template: '#chat-messages',
  ready: function () {
    this.$el.addEventListener("click", function (event) {
      event = event || win.event;

      if (event.target.tagName == "A") {
        win.open(event.target.href, "_blank");
        event.preventDefault();
        event.stopPropagation();
      }
    }, false);

    this.userScrolled = false;
    this.selfScroll = false;
    this.cont = this.cont || this.$el.querySelector(".chat-messages");
    this.timer = win.setInterval(this.scrollToBottom, 500);
  },
  methods: {
    onScroll: function () {
      if (this.selfScroll) {
        this.selfScroll = false;
        return;
      }

      var container = this.cont;
      this.userScrolled = container.scrollHeight - container.offsetHeight - container.scrollTop > 50;
    },

    scrollToBottom: function (e) {
      if (this.userScrolled) {
        return;
      }

      var container = this.cont;
      var loadedEventImage = e && e.loadedEventImage; // jshint ignore: line
      container.scrollTop = container.scrollHeight;
      this.selfScroll = true;
    }
  },
}));
