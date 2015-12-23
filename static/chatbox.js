  /*
  Copyright (c) 2015 Zohaib
  Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
  The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
  */

Zepto(function () {
  var s = io('http:///');
  var id = Date.now();
  var win = window;
  var $ = win.Zepto;
  var serverInfoGroup = 'SERVER';
  var currentGroup = serverInfoGroup;
  var txtarea = $('#dump');
  var msg = $('#msg');
  var maxhistory = $("#maxhistory");
  var sidebar = $(".sidebarContainer");
  var notificationSounds = {
    'notification': '/static/notif.mp3',
  };
  var MSG_DELIMETER = "~~~~>";
  var MSG_SENDER_REGEX = /^([^@]+)@+(.+)$/i;

  var getGroupHistoryLogger = function (name) {
    var logs = win.groupLogs = win.groupLogs || {};
    if (!logs[name]) {
      logs[name] = [];
      $(win).trigger("group-added", name);
    }

    return logs[name];
  };

  var setGroupHistoryLogger = function (name, log) {
    win.groupLogs = win.groupLogs || {};
    win.groupLogs[name] = log;
  };

  var parseMessage = function (data) {
      var msg = data.split(MSG_DELIMETER);
      if (msg.length < 2) {
        return null;
      }

      var matches = msg[0].match(MSG_SENDER_REGEX);
      if (matches && matches.length == 3){
        return {from: matches[1], to: matches[2], msg: msg[1]};
      }

      if (msg[0] != "SERVER") {
        return null;
      }

      return {from:msg[0], to: msg[0], msg:msg[1]};
  };

  var activateGroupPanel = function (newGroup) {
    if (txtarea.data("group") != newGroup) {
      txtarea.find(".incomingMessage").each(function (i, el) {
        el.remove();
      });

      var historyNodes = getGroupHistoryLogger(newGroup);
      $(historyNodes).each(function (i, el) {
        txtarea.prepend(el);
      });
      txtarea.data("group", newGroup);
    }
  };

  var writeGroupMessage = function (dataObj) {
    var msgTo = dataObj && dataObj.to;
    if (!msgTo) {
      console.warn("Invalid message", message);
      return;
    }

    var historyNodes = getGroupHistoryLogger(msgTo);
    var currentMsg = $(vxe.renderMessage('messageTemplate', dataObj));
    currentMsg.data("raw-data", dataObj);

    historyNodes.push(currentMsg);
    if (currentGroup == msgTo) {
      $(win).trigger("active-new-message", dataObj);
    }

    $(win).trigger("new-message", dataObj);
  };

  var sendMessage = function() {
    if (!s || !s.connected) {
      return;
    }

    var cmdResult = vxe.processComand(s, currentGroup, msg.val());

    if (!cmdResult && currentGroup != serverInfoGroup) {
      var m = {to: currentGroup, msg: msg.val()};
      if (m.msg.trim().length == 0){
        return;
      }

      s.emit("send-msg", m.to+MSG_DELIMETER+m.msg);
    }

    msg.val('');
    msg.focus();
  };

  var onGroupJoined = function(msg) {
    var joinInfo = msg.split('@');
    if (joinInfo && joinInfo.length >= 2) {
      var logger = getGroupHistoryLogger(joinInfo[1]);
      writeGroupMessage({from: serverInfoGroup, to: serverInfoGroup, msg: "Joined group "+joinInfo[1]});
      return;
    }

    console.error("Invalid join syntax", msg);
  };

  var onMessage = function(message) {
    var dataObj = parseMessage(message);
    writeGroupMessage(dataObj);
  };

  $("#btn").on("click", sendMessage);

  $("#dump, #bottomBar").on("click", function (){
    sidebar.hide();
  });

  $(".burgerButton").on("click", function(e) {
    sidebar.show();
    e.preventDefault();
    e.stopPropagation();
  });

  $(win).on("group-added", function (e, name) {
    var groupAnchor = $("<div class='groupName'><a href='#"+name+"'>"+name+"</a></div>");
    $("#groupList").append(groupAnchor);
    groupAnchor
      .data("name", name)
      .find("a")
      .on("click", function (){
        currentGroup = name;
        activateGroupPanel(name);
        e.preventDefault();
        e.stopPropagation();
      });
  });

  $(win).on("active-new-message", function (_, dataObj) {
    var name = dataObj.to;
    var historyNodes = getGroupHistoryLogger(name);

    // prepend the last message in the panel
    if (historyNodes.length && historyNodes[historyNodes.length - 1].parent().length == 0) {
      var currentMsg = historyNodes[historyNodes.length - 1];
      txtarea.prepend(currentMsg);
    }

    while (~~maxhistory.val() < historyNodes.length){
      var first = historyNodes[0];
      historyNodes = historyNodes.slice(1);
      setGroupHistoryLogger(name, historyNodes);
      first.remove();
    }
  });

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

  s.on('new-msg', onMessage);
  s.on('group-message', onMessage);
  s.on('group-join', onGroupJoined);
  s.on('connect', function() {
    s.on('disconnect', function(data) {
      writeGroupMessage({from: serverInfoGroup, to: serverInfoGroup, msg: "Disconnected..."});
    });

    writeGroupMessage({from: serverInfoGroup, to: serverInfoGroup, msg: "Connected..."});
  });

  writeGroupMessage({from: serverInfoGroup, to: serverInfoGroup, msg: "Connecting..."});
  msg.focus();
});
