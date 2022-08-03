const {defineConfig} = require("@vue/cli-service");
module.exports = defineConfig({
    pages: {
        index: {
            entry: "src/main.ts",
            title: "Lofalt"
        }
    },
    transpileDependencies: true,
    lintOnSave: false,
    configureWebpack: {},
    devServer: {
        open: true,
        port: 8080,
        host: "192.168.3.159",
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
