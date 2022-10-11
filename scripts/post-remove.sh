#!/bin/bash
for home in /home/*; do
    if [ -d "$home/.config/marketxiv" ]; then
        rm -rf "$home/.config/marketxiv"
    fi
done