import babel from 'rollup-plugin-babel';
import babelrc from 'babelrc-rollup';
import commonjs from 'rollup-plugin-commonjs';
import resolve from 'rollup-plugin-node-resolve';
import uglify from 'rollup-plugin-uglify';
import { minify } from 'uglify-js';

const plugins = [
    babel(babelrc()),
    resolve({
        jsnext: true
    }),
    commonjs()
];

if (process.env.COMPRESS) {
    plugins.push(uglify({}, minify));
}

let config = {
    input: 'client/main.js',
    output: {
        file: 'static/app.js',
        format: 'umd'
    },
    sourcemap: true,
    plugins: plugins
};

export default config;
