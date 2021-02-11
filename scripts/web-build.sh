#!/bin/sh
cd web
rm -rf app/* && npm run build
sass sass/main.sass styles/style.css