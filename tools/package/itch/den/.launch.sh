#!/bin/bash

if [ "$(uname)" == "Darwin" ]; then
    open -a Terminal -c "./osx/den $@"      
else
    x-terminal-emulator -e "./linux/den $@"
fi
