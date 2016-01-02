core.p2p = (function (win) {
  var moz = !!win.mozRTCPeerConnection;
  var PeerConnection = win.RTCPeerConnection || win.mozRTCPeerConnection || win.webkitRTCPeerConnection,
      SessionDescription = win.RTCSessionDescription || win.mozRTCSessionDescription || win.webkitRTCSessionDescription,
      IceCandidate = win.RTCIceCandidate || win.mozRTCIceCandidate || win.webkitRTCIceCandidate;

  var Peer2PeerDataConnection = function (opts) {
    $glueFunctions(this);
    this.options = opts || {channel: "WebRTCPeer", reliable: true};
    var servers = {
      iceServers: [
          {urls: "stun:23.21.150.121"},
          {urls: "stun:stun.l.google.com:19302"},
          {urls: "turn:numb.viagenie.ca", credential: "Hello123", username: " stun-register@yopmail.com"}
      ]
    };
    var pcConstraint = null;

    this.events = new core.EventEmitter();
    this.peerConnection = new PeerConnection(servers, pcConstraint);

    this.offers = [];
    this.iceCandidates = [];
    this.peerConnection.onicecandidate = this.onICECandidate;
    this.peerConnection.ondatachannel = this.onRecvDataChannel;

    var me = this;
  };

  Peer2PeerDataConnection.prototype = {
    close: function () {
      if (this.channel) {
        this.channel.close();
        this.channel = null;
      }

      if (this.peerConnection){
        this.peerConnection.close();
        this.offers = [];
        this.icecandidates = [];
        this.peerConnection = null;
      }
    },
    onICECandidate: function (e) {
      if (e.candidate) {
        this.iceCandidates.push(e.candidate);
        this.events.fire('candidate', e.candidate);
      }
    },
    addICECandidates: function (candidates) {
      for (var i = 0; i < candidates.length; i++) {
        console.log("Adding candidate", candidates[i]);
        this.peerConnection.addIceCandidate(new IceCandidate(candidates[i]));
      }
    },
    createOffer: function (cb) {
      var me = this;
      me.peerConnection.createOffer(
        function (descriptor) {
          me.peerConnection.setLocalDescription(descriptor);
          me.offers.push(descriptor);
          win.setTimeout(function () {
            cb && cb(descriptor);
          }, 0);
          me.events.fire("offer", descriptor);
          me.events.fire("offer.sdp", descriptor.sdp);
        }, function (e) {
          me.events.fire("offer.error", e);
        });
    },
    answerOffer: function (desc, cb) {
      var remoteDesc = new SessionDescription(desc);
      var me = this;
      me.peerConnection.setRemoteDescription(remoteDesc);
      me.peerConnection.createAnswer(
        function (descriptor) {
          me.peerConnection.setLocalDescription(descriptor);
          win.setTimeout(function () {
            cb && cb(descriptor);
          }, 0);
          me.events.fire("answer", descriptor);
          me.events.fire("answer.sdp", descriptor.sdp);
        }, function (e) {
          me.events.fire("answer.error", e);
        });
    },
    acceptAnswer: function (remoteDescriptor) {
      var remoteDesc = new SessionDescription(remoteDescriptor);
      var me = this;
      me.peerConnection.setRemoteDescription(remoteDesc);

    },
    createDataChannel: function () {
      this.channel = this.peerConnection.createDataChannel(
        this.options.channel || "WebRTCPeer",
        {
          reliable: (this.options.reliable || false)
        });

      this.hookDataChannelEvents();
    },
    onRecvDataChannel: function (event) {
      this.channel = event.channel;
      this.hookDataChannelEvents();
    },
    hookDataChannelEvents: function () {
      var me = this;
      var eventShooter = function (name) {
        return function (args) {
          me.events.fire(name, args);
        };
      };

      this.channel.onerror = eventShooter('error');
      this.channel.onmessage = eventShooter('message');
      this.channel.onopen = eventShooter('open');
      this.channel.onclose = eventShooter('close');
    },
  };

  return {
    Peer2PeerDataConnection: Peer2PeerDataConnection,
  };
})(window);
