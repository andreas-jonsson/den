#!/bin/bash

export BASE_VERSION=1.1.0
export FULL_VERSION=${BASE_VERSION}.0

rpl e34f19fc-289d-4fb9-b134-c1d07a29a273 "${BASE_VERSION}-0" snapcraft.yaml

go generate
snapcraft cleanbuild

snapcraft push --release=stable den_${BASE_VERSION}-0_amd64.snap
