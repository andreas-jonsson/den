#!/bin/bash

export BASE_VERSION=1.2.0

snapcraft cleanbuild
snapcraft push --release=stable den_${BASE_VERSION}-0_amd64.snap
