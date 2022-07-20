const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false,
  extends: ["plugin:vue/recommended", "eslint:recommended"],
});
