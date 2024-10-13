// import process from 'process'

// console.log(process.env)

// if (process.env.NODE_ENV !== 'development') {
document.write(
  '<script src="http://' +
    (location.host || 'localhost').split(':')[0] +
    ':35729/livereload.js"></' +
    'script>'
)
// }
