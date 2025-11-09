#!/bin/bash

set -e

echo "🔧 Building CCTV dashboard...."

go build -o bin/cctv-dashboard cmd/cctv-dashboard/main.go

