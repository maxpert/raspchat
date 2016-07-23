var gulp = require('gulp'),
    bower = require('main-bower-files'),
    pump = require('pump'),
    linker = require('gulp-merge-link'),
    rename = require('gulp-rename'),
    uglify = require('gulp-uglify'),
    concat = require('gulp-concat'),
    derequire = require('gulp-derequire'),
    htmlmin = require('gulp-htmlmin'),
    vueify = require('gulp-vueify');

var Settings = {
    assets: {
        source: [
        'static/index.html',
        'static/welcome.html',
        'static/*.svg',
        'static/*.png',
        'static/*.jpg',
        'static/*.gif',
        'static/*.css',
        'static/*.ttf',
        'static/favicon/**/*',
        'static/images/**/*'
        ],
        sourceBase: './static',
        output: 'dist/static'
    },
    js: {
        clientSourceFiles: [
            'static/js/core.js', 
            'static/js/rtc.js', 
            'static/js/peer_negotiator.js',
            'static/js/file_transfer.js',
            'static/js/components/*.js',
            'static/js/chat.js'
        ],
        clientFile: 'client.js',
        librariesFile: 'libraries.js',
        source: 'static',
        output: 'dist/static/js'
    },
    html: {
        merge: {
           '/static/js/client.min.js': [
                '/static/js/core.js', 
                '/static/js/components/*.js',
                '/static/js/chat.js'
           ],

           '/static/js/libraries.min.js': [
               '/static/bower_components/*/*.js',
               '/static/bower_components/**/*.js'
            ]
        },
        minify: {
            collapseWhitespace: true,
            conservativeCollapse: true,
        },
        output: 'dist/static'
    }
};

gulp.task('process-htmls', function () {
    return gulp.src(Settings.js.source + '/chat.html')
               .pipe(linker(Settings.html.merge))
               .pipe(htmlmin(Settings.html.minify))
               .pipe(rename('chat.html'))
               .pipe(gulp.dest(Settings.html.output));
});

gulp.task('bower-assemble', function() {
  return gulp.src(bower())
             .pipe(concat(Settings.js.librariesFile))
             .pipe(gulp.dest(Settings.js.output));
});

gulp.task('assemble', ['bower-assemble'], function () {
    return gulp.src(Settings.js.clientSourceFiles)
               .pipe(concat(Settings.js.clientFile))
               .pipe(gulp.dest(Settings.js.output));
});

gulp.task('copy-assets', function () {
    return gulp.src(Settings.assets.source, {base: Settings.assets.sourceBase})
               .pipe(gulp.dest(Settings.assets.output));
});

gulp.task('compile-assets', ['copy-assets', 'assemble', 'process-htmls'], function () {
    var librariesStream = gulp.src(Settings.js.output + '/' + Settings.js.librariesFile)
                              .pipe(uglify())
                              .pipe(rename({suffix: '.min'}))
                              .pipe(gulp.dest(Settings.js.output));

    var clientStream = gulp.src(Settings.js.output + '/' + Settings.js.clientFile)
                              .pipe(uglify())
                              .pipe(rename({suffix: '.min'}))
                              .pipe(gulp.dest(Settings.js.output));

    return pump([librariesStream, clientStream]);
});
