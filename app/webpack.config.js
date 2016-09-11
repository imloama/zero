// Learn more on how to config.
// - https://github.com/ant-tool/atool-build
var path = require("path");
var CopyWebpackPlugin = require('copy-webpack-plugin');

const webpack = require('atool-build/lib/webpack');


module.exports = function(webpackConfig, env) {
  webpackConfig.babel.plugins.push('transform-runtime');

  if (env === 'development') {
    webpackConfig.devtool = '#eval';
  }

  // CSS Modules Support.
  // Parse all less files as css module.
  webpackConfig.module.loaders.forEach(function(loader, index) {
    if (typeof loader.test === 'function' && loader.test.toString().indexOf('\\.less$') > -1) {
      loader.include = /node_modules/;
      loader.test = /\.less$/;
    }
    if (loader.test.toString() === '/\\.module\\.less$/') {
      loader.exclude = /node_modules/;
      loader.test = /\.less$/;
    }
    if (typeof loader.test === 'function' && loader.test.toString().indexOf('\\.css$') > -1) {
      loader.include = /node_modules/;
      loader.test = /\.css$/;
    }
    if (loader.test.toString() === '/\\.module\\.css$/') {
      loader.exclude = /node_modules/;
      loader.test = /\.css$/;
    }
  });

  webpackConfig.output.path = path.join(__dirname, '../public');

  //webpackConfig.context =  path.join(__dirname, 'app');

  //拷备静态文件
  var copyplugin =  new CopyWebpackPlugin([
          { from: path.join(__dirname, './index.html'), to: path.join(__dirname, '../public/index.html') }
        ]);
  webpackConfig.plugins.push(copyplugin);

  return webpackConfig;
};
