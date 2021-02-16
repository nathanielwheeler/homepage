const path = require('path')

module.exports = {
  entry: './js/main.js',
  mode: 'none',
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'app'),
  },
};