var path = require("path");

var webpack = require("webpack");

var CleanWebpackPlugin = require("clean-webpack-plugin");
var HtmlWebpackPlugin = require("html-webpack-plugin");

// Constants
const buildDir = path.join(__dirname, "dist");

// Configuration
module.exports = {
	mode: "development",
	devtool: "source-map",
	entry: [
		"./src/index.jsx"
	],
	output: {
		path: buildDir,
		filename: "[name].js"
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
		new CleanWebpackPlugin([buildDir]),
		new HtmlWebpackPlugin({
			title: "GH Gantt",
			meta: {
				viewport: "width=device-with, initial-scale=1"
			},
			template: "src/index.html"
		})
	]
};
