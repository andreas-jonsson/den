#!/bin/bash

if [ "$(uname)" == "Darwin" ]; then
    open -a Terminal -c "./osx/densrv $@"    
else
    x-terminal-emulator -e "./linux/densrv $@"
fi



