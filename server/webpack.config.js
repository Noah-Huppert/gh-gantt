var path = require("path");

var webpack = require("webpack");

const CleanObsoleteChunks = require('webpack-clean-obsolete-chunks');
var HtmlWebpackPlugin = require("html-webpack-plugin");
const ManifestPlugin = require("webpack-manifest-plugin");
const LiveReloadPlugin = require('webpack-livereload-plugin');
const CopyWebpackPlugin = require("copy-webpack-plugin");

// Configuration
module.exports = {
	mode: "development",
	devtool: "source-map",
	entry: [
		"./assets/js/index.jsx"
	],
	output: {
		filename: "[name].[hash].js",
		path: `${__dirname}/public`
	},
	module: {
		rules: [
			{
				test: /\.(js|jsx)/,
				exclude: /node_modules/,
				use: ["babel-loader"],
				resolve: {
					extensions: [".js", ".jsx"]
				}
			},
			{
				test: /\.scss/,
				use: ["style-loader", "css-loader", "sass-loader"]
			}
		]
	},
	plugins: [
		new CleanObsoleteChunks(),
		new HtmlWebpackPlugin({
			title: "GH Gantt",
			meta: {
				viewport: "width=device-with, initial-scale=1"
			},
			template: "assets/index.html"
		}),
		new ManifestPlugin({fileName: "manifest.json"}),
		new LiveReloadPlugin({appendScriptTag: true}),
		new CopyWebpackPlugin([{from: "./assets",to: ""}], {copyUnmodified: true,ignore: ["css/**", "js/**"] }),
	]
};
