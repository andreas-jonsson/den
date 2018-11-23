#!/bin/bash

curl -L -o butler.zip https://broth.itch.ovh/butler/linux-amd64/LATEST/archive/default
unzip butler.zip
chmod +x butler
./butler -V

./butler login
export USERVERSION=${BASE_VERSION}-${CI_PIPELINE_ID}

./butler push den-windows.zip phix/den:windows-edge --userversion $USERVERSION
./butler push den-osx.zip phix/den:osx-edge --userversion $USERVERSION
./butler push den-linux.zip phix/den:linux-edge --userversion $USERVERSION

./butler push densrv-windows.zip phix/den:server-windows-edge --userversion $USERVERSION
./butler push densrv-osx.zip phix/den:server-osx-edge --userversion $USERVERSION
./butler push densrv-linux.zip phix/den:server-linux-edge --userversion $USERVERSION
