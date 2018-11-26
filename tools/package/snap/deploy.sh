#!/bin/bash

echo $SNAPCRAFT_LOGIN_FILE | base64 --decode --ignore-garbage > ./../../../.snapcraft/snapcraft.cfg

snapcraft build
snapcraft push --release=edge den_git_amd64.snap