#!/bin/sh
cd web
# remove previous build
rm -rf app/* js/*
# Build javascript from typescript THEN build bundle
tsc 
# Bundle with webpack to app
npx webpack --config webpack.config.js
# Build styles from sass
sass sass/main.sass styles/style.css