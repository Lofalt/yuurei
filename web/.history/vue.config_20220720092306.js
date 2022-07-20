const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false,
  env: {
    browser: true,
    es2021: true,
    node: true,
  },
  extends: ["plugin:vue/essential", "airbnb-base"],
  parserOptions: {
    ecmaVersion: 12,
    parser: "@typescript-eslint/parser",
    sourceType: "module",
  },
  plugins: ["vue", "@typescript-eslint"],
  rules: {
    'indent': [
      'warn',
      'tab'
      //强制统一缩进
    ],
    'linebreak-style': [
      'error',
      'windows'
      //Windows的行结尾方式
    ],
    'quotes': [
      'warn',
      'single'
      //单引号
    ],
    'semi': [
      'warn',
      'never'
      //不加分号
    ]
  },
});
