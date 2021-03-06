const path = require('path');
const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const HtmlWebpackHarddiskPlugin = require('html-webpack-harddisk-plugin');
const HtmlWebpackInjectAttributesPlugin = require('html-webpack-inject-attributes-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

const publicPath = process.env.PUBLIC_PATH || '/static/';
const devPort = process.env.DEV_SERVER_PORT || 8081;
const devHost = process.env.DEV_SERVER_HOST || '0.0.0.0';
const appPort = process.env.APP_SERVER_PORT || 8080;

module.exports = (env, argv) => {
  const IS_DEV = argv.mode === 'development';

  const config = {
    output: {
      path: path.resolve(__dirname, './static'),
      publicPath: publicPath,
      filename: '[name].[hash].js',
      crossOriginLoading: 'anonymous',
    },
    devtool: 'source-map',
    devServer: {
      hot: true,
      port: devPort,
      host: devHost,
      publicPath: `http://localhost:${devPort}/`,
      headers: {
        'Access-Control-Allow-Origin': `http://localhost:${appPort}`
      }
    },
    module: {
      rules: [
        { test: /\.js$/, exclude: /node_modules/, use: 'babel-loader' },
        { test: /\.(sa|sc|c)ss$/, use: [(IS_DEV?'style-loader':MiniCssExtractPlugin.loader), 'css-loader', 'sass-loader'] },
        { test: /\.eot(\?v=.*)?$/, use: ['file-loader'] },
        { test: /\.(ico|png|gif|jpe?g)$/i, use: ['file-loader'] },
        { test: /\.woff2?(\?v=.*)?$/, use: [{ loader: 'url-loader', options: { prefix: 'font/', limit: 5000 } }] },
        { test: /\.ttf(\?v=.*)?$/, use: [{ loader: 'url-loader', options: { mimetype: 'application/octet-stream', limit: 10000 } }] },
        { test: /\.svg(\?v=.*)?$/, use: [{ loader: 'url-loader', options: { mimetype: 'image/svg+xml', limit: 10000 } }] }
      ]
    },
    resolve: {
      alias: {
        'react-dom': '@hot-loader/react-dom'
      }
    },
    plugins: [
      new HtmlWebpackPlugin({
        title: 'Seating',
        alwaysWriteToDisk: true,
      }),
      new HtmlWebpackHarddiskPlugin(),
      new HtmlWebpackInjectAttributesPlugin({
        crossorigin: 'anonymous',
      }),
      new webpack.DefinePlugin({
        MAPBOX_TOKEN: JSON.stringify(process.env.MAPBOX_TOKEN),
      })
    ]
  };

  if (IS_DEV) {
    config.output.publicPath = config.devServer.publicPath;
    config.plugins.push(new webpack.HotModuleReplacementPlugin());
  } else {
    config.plugins.push(new MiniCssExtractPlugin());
  }

  return config;
}