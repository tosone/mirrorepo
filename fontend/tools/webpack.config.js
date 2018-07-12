const path = require('path');
const webpack = require('webpack');
const AssetsPlugin = require('assets-webpack-plugin');

const pkg = require('../package.json');
const isDebug = global.DEBUG === false ? false : !process.argv.includes('--release');
const isVerbose = process.argv.includes('--verbose') || process.argv.includes('-v');
const useHMR = !!global.HMR;
const babelConfig = Object.assign({}, pkg.babel, {
  babelrc: false,
  cacheDirectory: useHMR,
  presets: pkg.babel.presets.map(x => (x === 'latest' ? ['latest', { es2015: { modules: false } }] : x)),
});

const config = {
  context: path.resolve(__dirname, '../src'),
  entry: ['!!style-loader!css-loader!react-mdl/extra/material.min.css', 'react-mdl/extra/material.min.js', './main.jsx'],
  output: {
    path: path.resolve(__dirname, '../public/dist'),
    publicPath: isDebug ? `http://localhost:${process.env.PORT || 3000}/dist/` : '/dist/',
    filename: isDebug ? '[name].js?[hash]' : '[name].[hash].js',
    chunkFilename: isDebug ? '[id].js?[chunkhash]' : '[id].[chunkhash].js',
    sourcePrefix: '  ',
  },
  devtool: isDebug ? 'source-map' : false,
  stats: {
    colors: true,
    reasons: isDebug,
    hash: isVerbose,
    version: isVerbose,
    timings: true,
    chunks: isVerbose,
    chunkModules: isVerbose,
    cached: isVerbose,
    cachedAssets: isVerbose,
  },
  plugins: [
    new webpack.DefinePlugin({
      'process.env.NODE_ENV': isDebug ? '"development"' : '"production"',
      __DEV__: isDebug,
    }),
    new AssetsPlugin({
      path: path.resolve(__dirname, '../public/dist'),
      filename: 'assets.json',
      prettyPrint: true,
    }),
    new webpack.LoaderOptionsPlugin({
      debug: isDebug,
      minimize: !isDebug,
    }),
  ],
  module: {
    rules: [
      {
        test: /\.jsx?$/,
        include: [path.resolve(__dirname, '../src'), path.resolve(__dirname, '../components')],
        loader: 'babel-loader',
        options: babelConfig,
      },
      {
        test: /\.css/,
        use: [
          {
            loader: 'style-loader',
          },
          {
            loader: 'css-loader',
            options: {
              sourceMap: isDebug,
              importLoaders: true,
              modules: true,
              localIdentName: isDebug ? '[name]_[local]_[hash:base64:3]' : '[hash:base64:4]',
              minimize: false,
            },
          },
        ],
      },
      {
        test: /\.json$/,
        exclude: [path.resolve(__dirname, '../src/routes.json')],
        loader: 'json-loader',
      },
      {
        test: /\.json$/,
        include: [path.resolve(__dirname, '../src/routes.json')],
        use: [
          {
            loader: 'babel-loader',
            options: babelConfig,
          },
          {
            loader: path.resolve(__dirname, './routes-loader.js'),
          },
        ],
      },
      {
        test: /\.md$/,
        loader: path.resolve(__dirname, './markdown-loader.js'),
      },
      {
        test: /\.(png|jpg|jpeg|gif|svg|woff|woff2)$/,
        loader: 'url-loader',
        options: {
          limit: 10000,
        },
      },
      {
        test: /\.(eot|ttf|wav|mp3)$/,
        loader: 'file-loader',
      },
    ],
  },
};
// Optimize the bundle in release (production) mode
if (!isDebug) {
  config.plugins.push(
    new webpack.optimize.UglifyJsPlugin({
      sourceMap: true,
      compress: {
        warnings: isVerbose,
      },
    }),
  );
  config.plugins.push(new webpack.optimize.AggressiveMergingPlugin());
}
// Hot Module Replacement (HMR) + React Hot Reload
if (isDebug && useHMR) {
  babelConfig.plugins.unshift('react-hot-loader/babel');
  config.entry.unshift('react-hot-loader/patch', 'webpack-hot-middleware/client');
  config.plugins.push(new webpack.HotModuleReplacementPlugin());
  config.plugins.push(new webpack.NoEmitOnErrorsPlugin());
}
module.exports = config;
