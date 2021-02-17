const path = require('path')

module.exports = {
  entry: {
    "main": './js/main.js'
  },
  mode: 'none',
  devtool: "source-map",
  output: {
    path: path.resolve(__dirname, 'app'),
    filename: '[name].js',
    sourceMapFilename: "[name].js.map"
  },
};