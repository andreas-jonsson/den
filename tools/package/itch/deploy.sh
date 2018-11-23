#!/bin/bash

curl -L -o butler.zip https://broth.itch.ovh/butler/linux-amd64/LATEST/archive/default
unzip butler.zip
chmod +x butler
./butler -V

./butler login
./butler push den.zip phix/den:edge --userversion ${BASE_VERSION}-${CI_PIPELINE_ID}
./butler push densrv.zip phix/den:edge --userversion ${BASE_VERSION}-${CI_PIPELINE_ID}
