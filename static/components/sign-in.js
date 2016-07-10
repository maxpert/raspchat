/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

(function (vue, win, doc, raspconfig) {
  var signInConfig = raspconfig.externalSignIn || {};

  var InvalidNickCharactersRegex = /[^a-zA-Z0-9]+/ig

  vue.component('google-sign-in', {
    template: '<div id="google-sign-in"></div>',

    ready: function () {
      if (!signInConfig.googleClientId) {
        return;
      }

      var head = doc.querySelector('head');
      var meta = doc.createElement('meta');
      meta.name = 'google-signin-client_id';
      meta.content = signInConfig.googleClientId;
      head.appendChild(meta);
      vue.nextTick(this.loadScript);
   },

   methods: {
     loadScript: function () {
       var funcName = '__google_sign_in_'+(new Date().getTime());
       var me = this;
       win[funcName] = this.scriptLoaded;
       var head = doc.querySelector('head');
       var script = doc.createElement('script');
       script.type='text/javascript';
       script.src='//apis.google.com/js/platform.js?onload='+funcName;
       head.appendChild(script);
     },

     scriptLoaded: function () {
       gapi.signin2.render('google-sign-in', {
        'scope': 'profile email',
        'width': 240,
        'height': 50,
        'longtitle': true,
        'theme': 'light',
        'onsuccess': this.onSuccess,
        'onfailure': this.onFailure
      });
     },

     onSuccess: function (user) {
       var profile = user.getBasicProfile();
       var userId = (profile.getEmail().split("@"))[0];
       var userInfo = {id: userId, name: profile.getName(), host: "google"};
       this.$dispatch("success", userInfo, user);
     },

     onFailure: function (err) {
       this.$dispatch("fail", err);
     }
   }

  });

  vue.component('login-form', {
    template: '#login-form',
    data: function () {
      return {
        isReady: false,
        isSignedIn: false,
        isValidNick: false,
        nick: '',
      };
    },

    ready: function () {
      this.$set('isReady', true);
      this.$watch('nick', this.onNickChanged);
      if (!raspconfig.externalSignIn) {
        this.$set('isSignedIn', true);
      }
    },

    methods: {
      googleSignInSuccess: function (userInfo) {
        localStorage["userInfo"] = JSON.stringify(userInfo);
        this.$set('isSignedIn', true);
        if (localStorage["userNick"]) {
          this.$set('nick', localStorage["userNick"]);
        } else {
          this.$set('nick', userInfo.id);
        }
      },

      onNickChanged: function () {
        if (this.nick.length > 0  && !this.nick.match(InvalidNickCharactersRegex)) {
          this.$set('isValidNick', true);
        }
        else {
          this.$set('isValidNick', false);
        }
      },

      signin: function () {
        if (!this.isValidNick) {
          return;
        }

        localStorage["userNick"] = this.nick;
        this.$dispatch('login', this.nick);
      }
    }
  });
})(Vue, window, window.document, window.RaspConfig);
