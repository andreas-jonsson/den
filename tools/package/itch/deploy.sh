#!/bin/bash

curl -L -o butler.zip https://broth.itch.ovh/butler/linux-amd64/LATEST/archive/default
unzip butler.zip
chmod +x butler
./butler -V

./butler login
export USERVERSION=${BASE_VERSION}-${CI_PIPELINE_ID}

./butler validate den
./butler push den phix/den:windows-osx-linux-edge --userversion $USERVERSION

./butler validate densrv
./butler push densrv phix/den:server-windows-osx-linux-edge --userversion $USERVERSION