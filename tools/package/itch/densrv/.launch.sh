#!/bin/bash

if [ "$(uname)" == "Darwin" ]; then
    open -a Terminal --args "./osx/densrv $@"    
else
    x-terminal-emulator -e "./linux/densrv $@"
fi
