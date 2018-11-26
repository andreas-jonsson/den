#!/bin/bash

# snapcraft login
# snapcraft export-login snapcraft.login
# base64 snapcraft.login | xsel --clipboard

echo $SNAPCRAFT_LOGIN_FILE | base64 --decode --ignore-garbage > ./../../../.snapcraft/snapcraft.cfg

rpl e34f19fc-289d-4fb9-b134-c1d07a29a273 "$BASE_VERSION-$CI_PIPELINE_ID" snapcraft.yaml

snapcraft build
