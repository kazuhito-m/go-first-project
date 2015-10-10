#!/bin/bash
# 依存性解決のgo get
go get github.com/bitly/go-simplejson
# テスト
go test ./...
if [[ $? -ne 0 ]] ; then
  echo 'Test failed and Build failed!'
  exit 1
fi
# ビルド・インストール
go install
