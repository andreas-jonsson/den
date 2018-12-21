#!/bin/bash

export BASE_VERSION=1.0.0
export FULL_VERSION=${BASE_VERSION}.0

go generate
tools/package/itch/build.sh

$BUTLER_DIR/butler validate den
$BUTLER_DIR/butler push den phix/den:windows-osx-linux --userversion $FULL_VERSION

$BUTLER_DIR/butler validate densrv
$BUTLER_DIR/butler push densrv phix/den:server-windows-osx-linux --userversion $FULL_VERSION