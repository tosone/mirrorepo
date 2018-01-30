const path = require('path');
const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const WebpackDevServer = require('webpack-dev-server');

const config = require('./config');

module.exports = {
  context: path.resolve(__dirname, '../src'),
  // entry: ['./main.jsx', 'webpack/hot/dev-server', 'webpack-dev-server/client?http://localhost:8080/'],
  entry: ['react-hot-loader/patch', 'webpack-hot-middleware/client', './main.jsx'],
  output: {
    path: path.join(__dirname, 'www'),
    filename: '[name].[hash].js',
    publicPath: '/',
  },
  devServer: {
    contentBase: './www',
    hot: true,
  },
  module: {
    rules: [
      {
        test: /\.jsx?$/,
        include: [path.resolve(__dirname, '../src'), path.resolve(__dirname, '../components')],
        loader: 'babel-loader',
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
              importLoaders: true,
              modules: true,
              localIdentName: '[name]_[local]_[hash:base64:3]',
              minimize: false,
            },
          },
        ],
      },
      {
        test: /\.md$/,
        use: [
          {
            loader: 'html-loader',
          },
          {
            loader: 'markdown-loader',
          },
        ],
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
      {
        test: /\.scss$/,
        use: [
          {
            loader: 'style-loader',
          },
          {
            loader: 'css-loader',
          },
          {
            loader: 'sass-loader',
          },
        ],
      },
    ],
  },
  plugins: [
    new webpack.NamedModulesPlugin(),
    new webpack.HotModuleReplacementPlugin(),
    // new webpack.optimize.UglifyJsPlugin({
    //   sourceMap: true,
    //   compress: {
    //     warnings: false,
    //   },
    // }),
    new HtmlWebpackPlugin({
      title: config.title,
      minify: { useShortDoctype: true },
      favicon: '../public/favicon.ico',
    }),
  ],
};
