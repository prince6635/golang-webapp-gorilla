var gulp = require('gulp');
var less = require('gulp-less');
var path = require('path');
var shell = require('gulp-shell');

var lessPath = './less/**/*.less';
var goPath = 'src/pistons/**/*.go';

gulp.task('less', function () {
  return gulp.src('less/app.less')
    .pipe(less({
      paths: [ path.join(__dirname, 'less', 'includes') ]
    }))
    .pipe(gulp.dest('./res/css'));
});

gulp.task('compilepkg', function() {
  return gulp.src(goPath, {read: false})
    .pipe(shell(['go install <%= stripPath(file.path) %>'],
      {
          templateData: {
            stripPath: function(path) {
              var subPath = path.substring(process.cwd().length + 5);
              var pkg = subPath.substring(0, subPath.lastIndexOf("\\"));
              return pkg;
            }
          }
      })
    );
});

gulp.task('watch', function() {
  gulp.watch(lessPath, ['less']);
  gulp.watch(goPath, ['compilepkg']);
});
