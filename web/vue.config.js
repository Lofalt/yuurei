const path = require('path');//引入path模块


// 导入compression-webpack-plugin
const CompressionWebpackPlugin = require('compression-webpack-plugin');
// 定义压缩文件类型
const productionGzipExtensions = ['js', 'css']

const {defineConfig} = require("@vue/cli-service");
module.exports = defineConfig({

    productionSourceMap: false,
    chainWebpack: (config) => {
        config.plugins.delete('prefetch')
        config.externals({
            'highlight.js': 'hljs',

        })
    },
    pages: {
        index: {
            entry: "src/main.ts",
            title: "幽居",
            favicon: "public/pratice.ico"
        }
    },
    transpileDependencies: true,
    lintOnSave: false,
    configureWebpack: {
        plugins: [
            new CompressionWebpackPlugin({
                filename: '[path][base].gz',
                algorithm: 'gzip',
                test: new RegExp('\\.(' + productionGzipExtensions.join('|') + ')$'),//匹配文件名
                threshold: 10240,//对10K以上的数据进行压缩
                minRatio: 0.8,
                deleteOriginalAssets: false,//是否删除源文件
            })
        ],
        module: {
            rules: [
                {
                    test: /\.(png|jpg|gif|jpeg)$/,
                    use: [{
                        loader: 'file-loader',
                        options: {
                            esModule: false,
                            name: "`images/[hash:8].[name].[ext]",
                        }
                    }],
                    type: 'javascript/auto',
                }

            ],
        },
    },
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
