#!/bin/sh

cd cmd/wasm
GOOS=js GOARCH=wasm go build -o ../../web/app/json.wasm