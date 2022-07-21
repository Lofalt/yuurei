const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false,
  dev: {
    open: true,
    port: 8080,
    '/api': {
      target: "localhost:8081"
    }
  }
});
