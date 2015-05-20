#!/usr/bin/env bash

echo ">> test 0.0.1"
go run *.go 0.0.1 http://localhost/tags.html 'v([-.\d]+).tar.gz'
echo

echo ">> test 1.0.0"
go run *.go 1.0.0 http://localhost/tags.html 'v([-.\d]+).tar.gz'
