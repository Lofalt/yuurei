const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false,
  devServer: {
    open: true,
    port: 8080,
    host: "192.168.3.173",
    proxy: {
      '/api': {
        target: "http://127.0.0.1:8081",
        pathRewrite: {
          '^/api': ''    //代理的路径
        }
      },
      '/img': {
        target: "http://127.0.0.1:8081"
      }
    }
  }
});
