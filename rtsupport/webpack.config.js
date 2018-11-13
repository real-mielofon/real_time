module.exports = {
	mode:"development",
	entry:'./index.js',
	output: {
		path: __dirname,
		filename: 'bundle.js'
	},
	devtool: 'inline-source-map',
	devServer: {
	  contentBase: './dist',
	  hot: true,
	  open: true
	},
	module: {
		rules:[
			{
				test: /\.(jsx|js)?$/,
				use: {
					loader: 'babel-loader',
					options: {
						presets: ["@babel/preset-env", "@babel/preset-react"]
					  }
				},
				exclude: /(node_modules)/,
			}
		]
	}
};