/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

var vue = require('vue');

vue.component('chat-message', vue.extend({
    props: ['message'],
    template: '#chat-message',
    ready: function () {
        this.hookImageLoads();
        this.$dispatch('chat-message-added', this.message);

        this.$watch('message.msg', function () {
            this.hookImageLoads();
            this.$dispatch('chat-message-added', this.message);
        }.bind(this));
    },
    methods: {
        imageLoaded: function (ev) {
            var me = this;
            me.$dispatch('chat-image-loaded', { loadedEventImage: ev });
        },
        hookImageLoads: function () {
            var imgs = this.$el.parentNode.querySelectorAll('img');
            for (var i in imgs) {
                var img = imgs[i];
                if (this._hasClass(img, 'avatar')) {
                    continue;
                }

                if (img.addEventListener) {
                    img.removeEventListener('load', this.imageLoaded);
                    img.addEventListener('load', this.imageLoaded, false);
                }
            }
        },

        _hasClass: function (element, selectorClass) {
            var idx = (' ' + element.className + ' ').replace(/[\n\t]/g, ' ').indexOf(selectorClass);
            return idx > -1;
        }
    }
}));
