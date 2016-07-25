/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

(function (vue, win, doc) {
  "use strict";

  function s4() {
    return Math.floor((1 + Math.random()) * 0x10000)
      .toString(16)
      .substring(1);
  }

  vue.component('files-upload', vue.extend({
    template: '#files-upload',
    props: {
    },
    data: function() {
      return {
        pendingFiles: [],
        hasPendingUploads: false,
        isDragActive: false
      };
    },
    ready: function () {
      if (!('FormData' in window)) {
        return;
      }

      doc.body.addEventListener("dragenter", this._activateDrag);
      doc.body.addEventListener("dragover", this._activateDrag);
      doc.body.addEventListener("dragleave", this._deactivateDrag);
      doc.body.addEventListener("dragend", this._deactivateDrag);
      doc.body.addEventListener("drop", this._deactivateDrag);

      doc.body.addEventListener("drop", this._droppedFiles);

      doc.body.addEventListener("drag", this._preventDragEvents);
      doc.body.addEventListener("dragstart", this._preventDragEvents);
      doc.body.addEventListener("dragend", this._preventDragEvents);
      doc.body.addEventListener("dragover", this._preventDragEvents);
      doc.body.addEventListener("dragenter", this._preventDragEvents);
      doc.body.addEventListener("dragleave", this._preventDragEvents);
      doc.body.addEventListener("drop", this._preventDragEvents);

      this._filesMap = {};
    },
    methods: {
      onFileStateChanged: function() {
      },

      onFileUploaded: function(fileInfo, response) {
        if (!this._filesMap[fileInfo.id]) {
          return;
        }

        delete this._filesMap[fileInfo.id];
        var foundIndex = this.pendingFiles.indexOf(fileInfo);
        if (foundIndex < 0) {
          return;
        }

        this.pendingFiles.splice(foundIndex, 1);
        if (response.failed === false && response.complete === true) {
          this._notifyUploaded(fileInfo, response);
        }
      },

      _notifyUploaded: function(fileInfo, uploadedInfo) {
        this.$dispatch('uploaded', fileInfo, uploadedInfo);
      },

      _activateDrag: function() {
        this.$set('isDragActive', true);
      },

      _deactivateDrag: function() {
        this.$set('isDragActive', false);
      },

      _droppedFiles: function(e) {
        var files = win.$arrayify(e.dataTransfer.files);
        var me = this;
        var selectedFiles = files.map(function(f) { 
          var fid = me._uuid();
          me._filesMap[fid] = f;

          return { 
            id: fid,
            file: f 
          };
        });

        this.pendingFiles = this.pendingFiles.concat(selectedFiles);
      },

      _preventDragEvents: function(e) {
        e.preventDefault();
        e.stopPropagation();
      },

      _uuid: function() {
        return s4() + s4() + '-' + s4() + '-' + s4() + '-' +
               s4() + '-' + s4() + s4() + s4();
      }
    }
  }));
})(window.Vue, window, window.document);
