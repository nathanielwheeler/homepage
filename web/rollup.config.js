import babel from 'rollup-plugin-babel'
import resolve from 'rollup-plugin-node-resolve';
import commonjs from 'rollup-plugin-commonjs';
import { eslint } from 'rollup-plugin-eslint';
import { createBasicConfig } from '@open-wc/building-rollup'
import pkg from './package.json';

const baseConfig = createBasicConfig()

export default {
  input: './js/main.js',
  plugins: [
    eslint(),
    resolve(),
    commonjs()
  ],
  output: {
    inlineDynamicImports: true,
    sourcemap: true,
    file: pkg.main,
    format: 'umd'
  }
}