const webpack = require('webpack');
const JavaScriptObfuscator = require('webpack-obfuscator');
const CopyPlugin = require("copy-webpack-plugin");

module.exports = {
    entry: './src/loader.js',
    output: {
        clean: true,
    },
    plugins: [
        new JavaScriptObfuscator ({
            stringArray: true,
            stringArrayEncoding: ['rc4'],
            rotateStringArray: true,
            //debugProtection: true,
            disableConsoleOutput: true,
            //selfDefending: true,
        }),
        new CopyPlugin({
            patterns: [
                { from: "src/style.css", to: "style.css" },
            ],
        }),
    ]
}