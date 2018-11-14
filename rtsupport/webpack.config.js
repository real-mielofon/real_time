var webpack = require('webpack');

module.exports = {
	mode: "development",
	entry: './index.js',
	output: {
		path: __dirname,
		filename: 'bundle.js'
	},
	devtool: 'inline-source-map',
	//	devtool: 'source-map'
	devServer: {
		contentBase: './',
		hot: true,
		open: true
	},
	module: {
		rules: [{
				test: /\.(jsx|js)?$/,
				use: {
					loader: 'babel-loader',
					options: {
						presets: ["@babel/preset-env", "@babel/preset-react"]
					}
				},
				exclude: /(node_modules)/,
			},
			{
				test: /\.(scss)$/,
				use: [{
					loader: 'style-loader', // inject CSS to page
				}, {
					loader: 'css-loader', // translates CSS into CommonJS modules
				}, {
					loader: 'postcss-loader', // Run post css actions
					options: {
						plugins: function () { // post css plugins, can be exported to postcss.config.js
							return [
								require('precss'),
								require('autoprefixer')
							];
						}
					}
				}, {
					loader: 'sass-loader' // compiles Sass to CSS
				}]
			},
			{
				test: /\.css$/,
				use: ['style-loader', 'css-loader']
			}
		]
	},
	plugins: [
		new webpack.ProvidePlugin({
			jQuery: 'jquery',
			$: 'jquery',
			jquery: 'jquery'
		})
	],
};