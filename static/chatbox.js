  /*
  Copyright (c) 2015 Zohaib
  Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
  The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
  */

Zepto(function () {
  var s = io('http:///');
  var id = Date.now();
  var channelName = "f00d3r";
  var $ = window.Zepto;

  var txtarea = $('#dump');
  var msg = $('#msg');
  var maxhistory = $("#maxhistory");
  var sidebar = $(".sidebarContainer");
  var channelMessages = {};

  $("#dump, #bottomBar").on("click", function (){
    sidebar.hide();
  });
  $(".burgerButton").on("click", function(e) {
    sidebar.show();
    e.preventDefault();
    e.stopPropagation();
  });

  var sendMessage = function() {
    if (!s || !s.connected) {
      return;
    }

    var cmdResult = vxe.processComand(s, channelName, msg.val());
    if (!cmdResult){
      var m = {to: channelName, msg: msg.val()};
      if (m.msg.trim().length == 0){
        return;
      }

      s.emit("send-msg", m.to+"~~~~>"+m.msg);
    }
    msg.val('');
    msg.focus();
  };

  $("#btn").on("click", sendMessage);

  msg.on("keydown", function (e) {
    var handled = false;
    if (e.which == 13) {
      if (!e.shiftKey) {
        sendMessage();
      }else{
        msg.val(msg.val() + "\n");
      }
      handled = true;
    }

    if (e.which == 9){
      msg.val(msg.val() + "\t");
      handled = true;
    }

    if (handled){
      e.preventDefault && e.preventDefault();
      e.stopPropagation();
      return false;
    }
  });

  var onMessage = function(data) {
    try {
      var msg = data.split("~~~~>");
      var dataObj = {from:msg[0], msg:msg[1]};

      // TODO: Maintain per channel history
      channelMessages[channelName] = channelMessages[channelName] || [];
      var nodes = channelMessages[channelName];
      var currentMsg =$(vxe.renderMessage('messageTemplate', dataObj));
      nodes.push(currentMsg);

      if (~~maxhistory.val() < nodes.length){
        var first = nodes[0];
        nodes = nodes.slice(1);
        first.remove();
      }

      txtarea.prepend(currentMsg);

      if (!$("#mute").prop("checked")) {
        var audio = new Audio('/static/notif.mp3');
        audio.play();
      }
    }
    catch(e) {
      txtarea.innerHTML += "\n<div>"+data+"</div>";
    }
  };

  s.on('new-msg', onMessage);
  s.on('group-message', onMessage);
  s.on('connect', function() {
    s.emit('join-group', channelName);
    s.on('disconnect', function(data) {
      vxe.renderMessage(txtarea, 'messageTemplate', {msg: "Disconnected..."});
      s.removeListener('.response', onMessage);
    });

    vxe.renderMessage(txtarea, 'messageTemplate', {msg: "Connected..."});
  });

  vxe.renderMessage(txtarea, 'messageTemplate', {msg: "Connecting..."});
  msg.focus();
  window.sock = s;
});
