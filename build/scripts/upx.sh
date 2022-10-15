#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $(basename $0) <file>"
    exit 1
fi

# If UPX is installed, compress the binary, otherwise do nothing.
if [ -x "$(command -v upx)" ]; then

    if [[ "$SKIP_COMPRESS" == "true" ]]; then
        echo "UPX: SKIP_COMPRESS is set to true, skipping compression."
        exit 0
    fi

    if [[ $1 == *"386"* ]]; then
        echo "UPX: Unable to compress 32-bit binaries."
        exit 0
    fi

    echo "UPX: Compressing $(basename $1)... (this may take a while)."
    timeout 20m upx $1 --best --ultra-brute
    
    # If the compression failed, skip it.
    if [ $? -eq 124 ]; then
        echo "UPX: Compression timed out for $(basename $1), skipping compression completely and cleaning up."
        cd $(dirname $1) && find . -type f -name "$(basename $1).*" ! -name "$(basename $1)" ! -name "$(basename $1).exe" -delete
        exit 0
    fi
else
    echo "UPX was not found, skipping compression."
fi