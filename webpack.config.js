const webpack = require('webpack');
const JavaScriptObfuscator = require('webpack-obfuscator');

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
        })
    ]
}