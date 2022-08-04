module.exports = {
    module: {
        rules: [
            {test: /\.css$/, use: 'css-loader'},
            {test: /\.ts$/, use: 'ts-loader'},
            {
                test: /\.(png|jpg|gif|jpeg)$/,
                use: [{
                    loader: 'file-loader',
                    options: {
                        esModule: false,
                        name:"`images/[hash:8].[name].[ext]",
                    }
                }],
                type: 'javascript/auto',
            }

        ],
    },
};