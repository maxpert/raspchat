/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

window.vxe = (function(win, doc) {
  var md = new markdownit();

  return {
    render: function (id, data) {
      var e = doc.getElementById(id);
      if (!e) {
        return '';
      }

      if (data.msg){
        data.msg = md.render(data.msg);
      }

      var tpl = (e.innerHTML || '').replace(/\s+/ig, ' ');
      var ret = Mustache.render(tpl, data);
      return ret;
    },

    renderMessage: function(template, data) {
      data.isMeta = !data.from;
      return "\n" + vxe.render(template, data);
    }
  };
})(window, window.document);

window.giffer = (function () {
  return {
    search: function (keywords, url_callback) {
      keywords = encodeURIComponent(keywords);
      var oReq = new XMLHttpRequest();
      oReq.addEventListener("load", function (r) {
        var resp = JSON.parse(oReq.response);
        console.log(resp);
        if (resp.attachments && resp.attachments.length > 0){
          var atch = resp.attachments[0];
          url_callback(atch.image_url, atch);
          return;
        }

        url_callback(null, null);
      });
      oReq.open("get", "/gif?q="+keywords);
      oReq.send();
    }
  };
})();

window.CommandProcessor = (function () {
  var nickRegex = /^\/nick\s+([A-Za-z0-9]+)$/i;
  var gifRegex = /^\/gif\s+(.+)$/i;

  return function (s, cmd) {
    var match = cmd.match(nickRegex);
    if (match) {
      s.emit("set-nick", match[1])
      return true;
    }

    match = cmd.match(gifRegex);
    if (match) {
      giffer.search(match[1], function (url, obj) {
        var t = cmd;
        if (url) {
          t = "!["+cmd+"]("+url+")";
        }

        s.emit("send-msg", channelName+"~~~~>"+t);
      });

      return true;
    }

    return false;
  }
})();
