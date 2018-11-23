#!/bin/bash

echo "Creating structure..."
export ROOT=.
rm -rf $ROOT/den-* $ROOT/densrv-*

cp -r $ROOT/tools/package/itch/den-* $ROOT
cp -r $ROOT/tools/package/itch/densrv-* $ROOT

export GOARCH=amd64

echo "Building Windows..."
export GOOS=windows
go build -o $ROOT/den-windows/den.exe $ROOT/den.go
go build -o $ROOT/densrv-windows/densrv.exe $ROOT/tools/densrv/densrv.go

echo "Building OSX..."
export GOOS=darwin
go build -o $ROOT/den-osx/den $ROOT/den.go
go build -o $ROOT/densrv-osx/densrv $ROOT/tools/densrv/densrv.go

echo "Building Linux..."
export GOOS=linux
go build -o $ROOT/den-linux/den $ROOT/den.go
go build -o $ROOT/densrv-linux/densrv $ROOT/tools/densrv/densrv.go

echo "Package files..."
rm -rf den-*.zip densrv-*.zip

zip -rq den-windows.zip $ROOT/den-windows
zip -rq den-osx.zip $ROOT/den-osx
zip -rq den-linux.zip $ROOT/den-linux

zip -rq densrv-windows.zip $ROOT/densrv-windows
zip -rq densrv-osx.zip $ROOT/densrv-osx
zip -rq densrv-linux.zip $ROOT/densrv-linux