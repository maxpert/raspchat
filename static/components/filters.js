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

  vue.filter('falsy_to_block_display', function (value) {
    return value ? 'block' : 'none';
  });


  var fragmentNode = document.createDocumentFragment();
  var virtualDiv = document.createElement('div');
  fragmentNode.appendChild(virtualDiv);
  vue.filter('emojify', function (value) {
    if (!emojify) {
      return value;
    }

    virtualDiv.innerHTML = value;
    emojify.run(virtualDiv);
    return virtualDiv.innerHTML;
  });

  vue.filter('avatar_url', function (value) {
    // return '//api.adorable.io/avatars/face/eyes6/nose7/face1/AA0000';
    return '//api.adorable.io/avatars/256/zmg-' + value + '.png';
  });
})(Vue, window, window.document);
