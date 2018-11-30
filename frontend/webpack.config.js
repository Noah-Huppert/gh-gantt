var path = require("path")

var webpack = require("webpack")

const CleanWebpackPlugin = require("clean-webpack-plugin")
const HtmlWebpackPlugin = require("html-webpack-plugin")
const ExtractTextPlugin = require("extract-text-webpack-plugin")

// Constants
const buildDir = path.join(__dirname, "dist")

// Configuration
module.exports = {
	mode: "development",
	devtool: "inline-source-map",
	//devtool: "source-map",
	entry: [
		"./src/index.js"
	],
	output: {
		path: buildDir,
		filename: "[name].js"
	},
	resolve: {
		alias: {
			"vue$": "vue/dist/vue.js",
			//"vue-router$": "vue-router/dist/vue-router.js"
		},
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
				test: /\.s(a|c)ss/,
				use: ExtractTextPlugin.extract({
					fallback: "style-loader",
					use: ["css-loader", "sass-loader"]
				})
			},
			{
				test: /\.(woff|woff2|eot|ttf|svg)$/,
				loader: ["file-loader"]
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
		}),
		new ExtractTextPlugin("css/styles.css")
	]
}
