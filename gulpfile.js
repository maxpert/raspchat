var gulp = require('gulp'),
    rename = require('gulp-rename'),
    uglify = require('gulp-uglify'),
    htmlmin = require('gulp-htmlmin'),
    rev = require('gulp-rev-append');

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
        sourceFile: 'chat.js',
        outputFile: 'app.js',
        static: 'static',
        source: 'static/js',
        output: 'dist/static',
        server: {
            src: 'src/*.js',
            dest: 'dist/server'
        }
    },
    html: {
        minify: {
            collapseWhitespace: true,
            conservativeCollapse: true,
        },
        source: 'static',
        output: 'dist/static'
    }
};

gulp.task('process-htmls', function () {
    return gulp.src( Settings.html.source + '/chat.html')
        .pipe(rev())
        .pipe(htmlmin(Settings.html.minify))
        .pipe(rename('chat.html'))
        .pipe(gulp.dest(Settings.html.output));
});

gulp.task('copy-assets', function () {
    return gulp.src(Settings.assets.source, {base: Settings.assets.sourceBase})
        .pipe(gulp.dest(Settings.assets.output));
});

gulp.task('compile-assets', ['copy-assets', 'process-htmls'], function () {
    return gulp.src(Settings.js.static + '/' + Settings.js.outputFile)
        .pipe(uglify())
        .pipe(rename(Settings.js.outputFile))
        .pipe(gulp.dest(Settings.js.output));
});

gulp.task('dist-package', function() {
    return gulp.src('package*.json')
        .pipe(gulp.dest('dist'));
});

gulp.task('dist-src', function() {
    return gulp.src(Settings.js.server.src)
        .pipe(gulp.dest(Settings.js.server.dest));
});

gulp.task('default', ['compile-assets', 'dist-src', 'dist-package']);
