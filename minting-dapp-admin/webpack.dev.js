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
<<<<<<< HEAD
=======
    proxy: {
      '/api': {
        target: 'http://localhost:4040',
        router: () => 'http://localhost:3999',
      }
    },
>>>>>>> 910ce0f57929e011abf7ac9af44dd93f7c3ef56c
  },
};
