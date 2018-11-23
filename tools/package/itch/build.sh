#!/bin/bash

echo "Creating structure..."
export ROOT=.
rm -rf $ROOT/den $ROOT/densrv

mkdir $ROOT/den
mkdir $ROOT/den/windows
mkdir $ROOT/den/osx
mkdir $ROOT/den/linux

cp $ROOT/tools/package/itch/den.itch.toml $ROOT/den/.itch.toml 

mkdir $ROOT/densrv
mkdir $ROOT/densrv/windows
mkdir $ROOT/densrv/osx
mkdir $ROOT/densrv/linux

cp $ROOT/tools/package/itch/densrv.itch.toml $ROOT/densrv/.itch.toml 

export GOARCH=amd64

echo "Building Windows..."
export GOOS=windows
go build -o $ROOT/den/windows/den.exe $ROOT/den.go
go build -o $ROOT/densrv/windows/densrv.exe $ROOT/tools/densrv/densrv.go

echo "Building OSX..."
export GOOS=darwin
go build -o $ROOT/den/osx/den $ROOT/den.go
go build -o $ROOT/densrv/osx/densrv $ROOT/tools/densrv/densrv.go

echo "Building Linux..."
export GOOS=linux
go build -o $ROOT/den/linux/den $ROOT/den.go
go build -o $ROOT/densrv/linux/densrv $ROOT/tools/densrv/densrv.go

echo "Package files..."
rm -rf den.zip densrv.zip
zip -rq den.zip $ROOT/den
zip -rq densrv.zip $ROOT/densrv
