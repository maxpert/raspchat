var gulp = require('gulp'),
    pump = require('pump'),
    rename = require('gulp-rename'),
    uglify = require('gulp-uglify'),
    concat = require('gulp-concat'),
    derequire = require('gulp-derequire'),
    browserify = require('gulp-browserify'),
    htmlmin = require('gulp-htmlmin'),
    argv = require('yargs').alias('d', 'debug').argv;

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
        output: 'dist/static'
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
               .pipe(htmlmin(Settings.html.minify))
               .pipe(rename('chat.html'))
               .pipe(gulp.dest(Settings.html.output));
});

gulp.task('compile-js', function () {
    return gulp.src(Settings.js.source + '/' + Settings.js.sourceFile)
        .pipe(browserify({
            insertGlobals : true,
            debug : !argv.p
        }))
        .pipe(rename(Settings.js.outputFile))
        .pipe(gulp.dest(Settings.js.static));
});

gulp.task('copy-assets', function () {
    return gulp.src(Settings.assets.source, {base: Settings.assets.sourceBase})
               .pipe(gulp.dest(Settings.assets.output));
});

gulp.task('compile-assets', ['copy-assets', 'compile-js', 'process-htmls'], function () {
    return gulp.src(Settings.js.static + '/' + Settings.js.outputFile)
                .pipe(uglify())
                .pipe(rename(Settings.js.outputFile))
                .pipe(gulp.dest(Settings.js.output));
});
