const path = require('path');

module.exports = {
	mode: 'development',
	devtool: 'source-map',
	devServer: {
		static: {
			directory: path.join(__dirname, 'public/'),
		},
		compress: true,
		port: 4040,
		proxy: {
			'/api': {
				target: 'http://localhost:4040',
				router: () => 'http://localhost:3999',
			}
		},
	},
};