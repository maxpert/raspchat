(function (vue, win, doc) {
  var md = new markdownit("default", {
    linkify: true,
  });

  vue.filter('markdown', function (value) {
    return md.render(value);
  });

  vue.filter('better_date', function (value) {
    return moment(value).calendar();
  });

  vue.filter('escape_html', function (value) {
    return he.encode(value);
  });

  vue.filter('avatar_url', function (value) {
    // return 'http://api.adorable.io/avatars/face/eyes6/nose7/face1/AA0000';
    return 'http://api.adorable.io/avatars/256/zmg-' + value + '.png';
  });
})(Vue, window, window.document);