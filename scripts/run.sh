#!/bin/sh
./scripts/sass.sh
./scripts/wasm.sh
go run cmd/homepage/main.go