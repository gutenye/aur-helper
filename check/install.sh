#!/usr/bin/env bash

mkdir -p dist
go build -o dist/gutaur-check *.go
ln -sf `pwd`/dist/gutaur-check ~/bin
