const path = require('path');
const { merge } = require('webpack-merge');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const TsConfigPathsPlugin = require('tsconfig-paths-webpack-plugin');
// remove ? (275kb/80kb gzipped)
const { BundleStatsWebpackPlugin } = require('bundle-stats-webpack-plugin');

const webpackDev = require('./webpack.dev');

module.exports = merge([
  {
    mode: 'production',
    entry: './src/index.tsx',
    output: {
      filename: '[name].[contenthash].js',
      path: path.resolve(__dirname, 'dist/'),
      clean: true,
    },
    module: {
      rules: [
        {
          test: /\.(ts|tsx)$/,
          exclude: /node_modules/,
          resolve: {
            extensions: ['.ts', '.tsx', '.json'],
            plugins: [
              new TsConfigPathsPlugin(),
            ],
          },
          use: 'ts-loader',
        },
        {
          test: /\.scss$/,
          use: [
            { loader: MiniCssExtractPlugin.loader },
            { loader: "css-loader", options: { modules: true } },
            { loader: "sass-loader" },
          ],
        },
      ],
    },
    plugins: [
      new HtmlWebpackPlugin({
        template: './public/index.html',
      }),
      new MiniCssExtractPlugin({
        filename: '[name].[contenthash].css',
      }),
      new BundleStatsWebpackPlugin(),
    ],
  },
  process.env.NODE_ENV !== 'production' ? webpackDev : {},
]);
