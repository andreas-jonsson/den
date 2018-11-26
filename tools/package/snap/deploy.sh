#!/bin/bash

# snapcraft login
# snapcraft export-login snapcraft.login
# base64 snapcraft.login | xsel --clipboard

echo $SNAPCRAFT_LOGIN_FILE | base64 --decode --ignore-garbage > ./../../../.snapcraft/snapcraft.cfg

snapcraft build
snapcraft push --release=edge den_git_amd64.snap