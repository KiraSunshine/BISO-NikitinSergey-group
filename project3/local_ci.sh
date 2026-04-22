#!/usr/bin/env bash
set -euo pipefail

echo "[1/5] gofmt"
gofmt -w ./cmd ./internal

echo "[2/5] unit tests"
go test -v -race -covermode=atomic -coverprofile=coverage.out ./...

echo "[3/5] install golangci-lint"
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8

echo "[4/5] install gosec"
go install github.com/securego/gosec/v2/cmd/gosec@v2.22.4

echo "[5/5] lint + sast + build"
"$(go env GOPATH)/bin/golangci-lint" run ./...
"$(go env GOPATH)/bin/gosec" ./...
mkdir -p dist
go build -o dist/app ./cmd/app

echo
echo "OK: local CI passed"
