#!/bin/bash

export CCTV_CAMERA=rtsp://CCTV_IP:CCTV_PORT/live/ch00_0
export RTSP_SERVER=rtsp://CCTV_SERVER_IP:CCTV_SERVER_PORT/cctv_01

go run cmd/cctv-publisher/main.go