/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

var vue = require('vue');
var utils = require('../vendor/utils');

var uploadViaXHR = function (file, progressCallback, successCallback, failCallback) {
    var fd = new FormData();
    fd.append('file', file);
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/file', true);

    xhr.upload.onprogress = function (e) {
        if (e.lengthComputable) {
            var done = e.position || e.loaded, total = e.totalSize || e.total;
            progressCallback && progressCallback(done, total); // jshint ignore: line
        }
    };

    xhr.onload = function () {
        if (this.status == 200) {
            successCallback && successCallback(xhr); // jshint ignore: line
        }
        else {
            failCallback && failCallback(xhr.status, xhr); // jshint ignore: line
        }
    };

    xhr.send(fd);
};

vue.component('file-uploader', vue.extend({
    template: '#file-uploader',
    props: {
        pending: {
            type: Boolean
        },

        failed: {
            type: Boolean
        },

        complete: {
            type: Boolean
        },

        fileInfo: {
            required: true
        }
    },
    data: function () {
        return {
            totalBytes: 0,
            uploadedBytes: 0
        };
    },
    computed: {
        percentComplete: function () {
            if (this.totalBytes === 0 || this.uploadedBytes === 0) {
                return 0;
            }

            return Math.floor(this.uploadedBytes * 100.0 / this.totalBytes) + '';
        }
    },
    ready: function () {
        this.$set('pending', true);
        this.$set('complete', false);
        this.$set('failed', true);

        vue.nextTick(this._startUpload);
    },

    methods: {
        _notifyStatusChanged: function (data) {
            var eventParams = utils.Mix({
                pending: this.pending,
                failed: this.failed,
                complete: this.complete
            }, data);

            this.$dispatch('state-changed', this.fileInfo, eventParams);
            if (this.complete && data) {
                this.$dispatch('file-uploaded', this.fileInfo, eventParams);
            }
        },

        _startUpload: function () {
            this.$set('pending', true);
            uploadViaXHR(this.fileInfo.file, this._uploadProgress, this._uploadSuccess, this._uploadError);
        },

        _uploadProgress: function (uploaded, total) {
            this.$set('pending', true);
            this.$set('complete', false);
            this.$set('failed', false);
            this.$set('totalBytes', total);
            this.$set('uploadedBytes', uploaded);
            this._notifyStatusChanged({});
        },

        _uploadSuccess: function (xhr) {
            this.$set('pending', false);
            this.$set('complete', true);
            this.$set('failed', false);
            this._notifyStatusChanged(JSON.parse(xhr.responseText));
        },

        _uploadError: function (errorCode, xhr) {
            this.$set('pending', false);
            this.$set('complete', true);
            this.$set('failed', true);
            this._notifyStatusChanged({ errorCode: errorCode, error: xhr.responseText });
        }
    }
}));
