#!/bin/bash

# check if folder ./stream exists, if not, create it
if [ ! -d "./stream" ]; then
    mkdir ./stream
fi

cp public/index.html ./stream

go build -o ffmpeg-cctv-reader
