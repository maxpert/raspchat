/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

var vue = require('vue');

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
            if (e.shiftKey) {
                this.$set('message', msg + '\n');
                return;
            }

            this.$set('message', '');
            this.$dispatch('send-message', msg);
            this.$el.querySelector('.msg').focus();
        },

        tabPressed: function () {
            var msg = this.$get('message');
            this.$set('message', msg + '  ');
        },

        onFileUploaded: function (fileInfo, uploadInfo) {
            var message = '[' + fileInfo.file.name + '](' + uploadInfo.url + ')';
            if (fileInfo.file.type.startsWith('image/')) {
                message = '!' + message + '\n\n **IMAGE** ' + message;
            } else if (fileInfo.file.type.startsWith('video/')) {
                message = '!' + message + '\n\n **VIDEO** ' + message;
            } else {
                message = '**FILE** ' + message;
            }

            this.$dispatch('send-message', message);
            this.$el.querySelector('.msg').focus();
        },

        onFileUploadFailed: function (fileInfo) {
            window.alert('Unable to upload file ' + fileInfo.file.name);
            this.$el.querySelector('.msg').focus();
        }
    },
}));
