#!/bin/bash

echo "Creating structure..."
rm -rf den densrv

cp -r tools/package/itch/den .
mkdir den/windows
mkdir den/osx
mkdir den/linux

cp -r tools/package/itch/densrv .
mkdir densrv/windows
mkdir densrv/osx
mkdir densrv/linux

export GOARCH=amd64

echo "Building Windows..."
export GOOS=windows
go build -o den/windows/den.exe den.go
go build -o densrv/windows/densrv.exe tools/densrv/densrv.go

echo "Building OSX..."
export GOOS=darwin
go build -o den/osx/den den.go
go build -o densrv/osx/densrv tools/densrv/densrv.go

echo "Building Linux..."
export GOOS=linux
go build -o den/linux/den den.go
go build -o densrv/linux/densrv tools/densrv/densrv.go
