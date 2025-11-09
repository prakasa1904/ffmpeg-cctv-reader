#!/bin/bash

set -e

echo "🔧 Building FFMPEG video forwarder...."

go build -o bin/cctv-forwarder cmd/cctv-forwarder/main.go
