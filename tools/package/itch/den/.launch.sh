#!/bin/bash

if [ "$(uname)" == "Darwin" ]; then
    open -a Terminal --args "./osx/den $@"      
else
    x-terminal-emulator -e "./linux/den $@"
fi
