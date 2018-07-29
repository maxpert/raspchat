/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

var markdownit = require('markdown-it');
var markdownitHTML5Embed = require('markdown-it-html5-embed');
var moment = require('moment');
var he = require('he');
var vue = require('vue');
var emojify = require('emojify.js');

var md = new markdownit('default', {
    linkify: true
}).use(markdownitHTML5Embed);

vue.filter('markdown', function (value) {
    window.debug && window.debug.filters && console.log(value);
    return md.render(value);
});

vue.filter('better_date', function (value) {
    window.debug && window.debug.filters && console.log(value);
    return moment(value).calendar();
});

vue.filter('escape_html', function (value) {
    window.debug && window.debug.filters && console.log(value);
    return he.encode(value);
});

vue.filter('falsy_to_block_display', function (value) {
    window.debug && window.debug.filters && console.log(value);
    return value ? 'block' : 'none';
});

vue.filter('friendly_progress', function (value) {
    window.debug && window.debug.filters && console.log(value);
    if (~~value >= 100) {
        return 'almost done...';
    }

    return 'uploaded ' + value + '%';
});


var fragmentNode = document.createDocumentFragment();
var virtualDiv = document.createElement('div');
fragmentNode.appendChild(virtualDiv);
vue.filter('emojify', function (value) {
    virtualDiv.innerHTML = value;
    emojify.run(virtualDiv);
    return virtualDiv.innerHTML;
});

function hashFnv32a(str, seed) {
    /*jshint bitwise:false */
    var i, l,
        hval = (seed === undefined) ? 0x811c9dc5 : seed;

    for (i = 0, l = str.length; i < l; i++) {
        hval ^= str.charCodeAt(i);
        hval += (hval << 1) + (hval << 4) + (hval << 7) + (hval << 8) + (hval << 24);
    }

    return hval >>> 0;
}

vue.filter('avatar_url', function (value) {
    var type = hashFnv32a(value) % 2 ? 'female' : 'male';
    return '//avatars.dicebear.com/v1/'+type+'/'+value+'/128.png';
});
