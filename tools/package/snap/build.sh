#!/bin/bash

# snapcraft login
# snapcraft export-login snapcraft.login
# base64 snapcraft.login | xsel --clipboard

echo $SNAPCRAFT_LOGIN_FILE | base64 --decode --ignore-garbage > ./../../../.snapcraft/snapcraft.cfg

export CURRENT_VERSION=1.2.0-0
rpl ${CURRENT_VERSION} "${BASE_VERSION}-${CI_PIPELINE_ID}" snapcraft.yaml

go generate
snapcraft build
