#!/bin/bash

export BASE_VERSION=1.1.0
export FULL_VERSION=${BASE_VERSION}.0

go generate
snapcraft cleanbuild

snapcraft push --release=stable den_${BASE_VERSION}-0_amd64.snap
