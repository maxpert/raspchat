window.core.PeerConnectionNegotiator = (function (win) {

  var installPeerDebugHooks = function (peerConnection) {
      peerConnection.events.on('offer', function (desc) {
        console.log("OFFER", desc);
      });

      peerConnection.events.on("candidate", function (c) {
        console.log("CANDIDATE", c);
      });

      peerConnection.events.on("open", function () {
        console.log("OPENED", peerConnection.channel);
      });

      peerConnection.events.on("close", function () {
        console.log("CLOSED", peerConnection.channel);
      });

      peerConnection.events.on("message", function (msg) {
        console.log("MESSAGE", msg);
      });
  };

  var PeerConnectionNegotiator = function (transport, options) {
    $glueFunctions(this);
    this.events = new core.EventEmitter();
    this.dataBuffer = [];
    this.channel = null;

    this.peerConnection = new core.p2p.Peer2PeerDataConnection(options);
    this.peerConnection.events.on("open", this.onPCOpen);
    this.peerConnection.events.on("close", this.onPCClose);
    this.peerConnection.events.on("error", this.onPCError);
    this.peerConnection.events.on("data", this.onPCData);

    this.transport = transport;
    this.transport.events.on("raw-message", this.onOfferRecv);
    this.transport.events.on("raw-message", this.onAnswerRecv);

    installPeerDebugHooks(this.peerConnection);
  };

  PeerConnectionNegotiator.prototype = {
    close: function () {
      this.transport.events.off("raw-message", this.onOfferRecv);
      this.transport.events.off("raw-message", this.onAnswerRecv);
      this.peerConnection.close();

      this.channel = null;
      this.peerConnection = null;
      this.transport = null;
    },

    onPCOpen: function () {
      for(var i = 0; i < this.dataBuffer.length; i++){
        this.channel.send(this.dataBuffer[i]);
      }

      this.channel = this.peerConnection.channel;
      this.events.fire("ready", this.peerConnection.channel);
    },

    onPCClose: function () {
      var oldCh = this.channel;
      this.channel = null;
      this.events.fire("close", oldCh);
    },

    onPCError: function (err) {
      this.events.fire("error", err);
    },

    onPCData: function (event) {
      this.events.fire("data", event.data, event);
    },

    send: function (data) {
      if (!this.channel) {
        this.dataBuffer.push(data);
        return;
      }

      this.channel.send(data);
    },

    onOfferRecv: function (from, reqMsg) {
      if (reqMsg.type != "P2PHandShake" || reqMsg.mode != "RequestOffer") {
        return;
      }

      console.log("RequestOffer", reqMsg);
      var me = this;
      me.peerConnection.answerOffer(reqMsg.offer, function (offer) {
        if (reqMsg.candidates) {
          me.peerConnection.addICECandidates(reqMsg.candidates);
          win.setTimeout(function () {
            me.transport.sendRaw(from, {
              type: "P2PHandShake",
              mode: "OfferResponse",
              offer: offer,
              candidates: me.peerConnection.iceCandidates});
            console.log("OfferedResponse", reqMsg);
          }, 3000);
        }
      });
    },

    onAnswerRecv: function (from, reqMsg) {
      if (reqMsg.type != "P2PHandShake" || reqMsg.mode != "OfferResponse") {
        return;
      }

      console.log("OfferResponse", reqMsg);
      this.peerConnection.acceptAnswer(reqMsg.offer);
      if (reqMsg.candidates) {
        this.peerConnection.addICECandidates(reqMsg.candidates);
      }
    },

    connectTo: function (to) {
      var me = this;
      me.peerConnection.createDataChannel();
      me.peerConnection.createOffer(function (offer) {
        win.setTimeout(function () {
          me.transport.sendRaw(to, {
            type: "P2PHandShake",
            mode: "RequestOffer",
            offer: offer,
            candidates: me.peerConnection.iceCandidates
          });
        }, 2000);
      });
    },
  };

  return PeerConnectionNegotiator;
})(window);
