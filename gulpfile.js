var gulp = require('gulp'),
    merge = require('merge-stream'),
    bower = require('main-bower-files'),
    combiner = require('stream-combiner2'),
    rename = require('gulp-rename'),
    uglify = require('gulp-uglify'),
    concat = require('gulp-concat'),
    vueify = require('gulp-vueify');

var Settings = {
    js: {
        clientSourceFiles: [
            'static/rtc.js', 
            'static/core.js', 
            'static/peer_negotiator.js',
            'static/file_transfer.js',
            'static/components/*.js'
        ],
        clientFile: 'client.js',
        librariesFile: 'libraries.js',
        source: 'static',
        output: 'dist/static/js'
    }
};

function combine(streams) {
  var combined = combiner.obj(streams);

  // any errors in the above streams will get caught
  // by this listener, instead of being thrown:
  combined.on('error', console.error.bind(console));
  return combined;
}

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

gulp.task('build', ['assemble', 'bower-assemble'], function () {
    var librariesStream = gulp.src(Settings.js.output + '/' + Settings.js.librariesFile)
                              .pipe(uglify())
                              .pipe(rename({suffix: '.min'}))
                              .pipe(gulp.dest(Settings.js.output));

    var clientStream = gulp.src(Settings.js.output + '/' + Settings.js.clientFile)
                              .pipe(uglify())
                              .pipe(rename({suffix: '.min'}))
                              .pipe(gulp.dest(Settings.js.output));

    return combine(librariesStream, clientStream);
});
