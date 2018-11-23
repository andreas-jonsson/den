#!/bin/bash

curl -L -o butler.zip https://broth.itch.ovh/butler/linux-amd64/LATEST/archive/default
unzip butler.zip
chmod +x butler
./butler -V

export USERVERSION=${BASE_VERSION}-${CI_PIPELINE_ID}

./butler login
./butler push den.zip phix/den:windows-osx-linux-edge --userversion $USERVERSION
./butler push densrv.zip phix/den:windows-osx-linux-server-edge --userversion $USERVERSION
