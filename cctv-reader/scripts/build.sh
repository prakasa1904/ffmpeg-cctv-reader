#!/bin/bash

set -e

echo "🔧 Build FFMPEG video publisher...."

go build -o ffmpeg-forwarder cctv-reader/main.go