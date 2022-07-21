const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false,
  server: {
    open: true,
    port: 8080,
    Proxy: {
      '/api': {
        target: "localhost:8081"
      }
    }
  }
});
