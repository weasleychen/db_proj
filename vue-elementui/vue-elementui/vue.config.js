module.exports = {
  devServer: {
    proxy: {
      '/api': {
        target: 'http://8.138.93.57:10000',
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        }
      }
    }
  }
}
