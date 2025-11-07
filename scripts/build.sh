#!/bin/bash

# check if folder ./stream exists, if not, create it
if [ ! -d "./stream" ]; then
    mkdir ./stream
fi

go build -o ffmpeg-cctv-reader
