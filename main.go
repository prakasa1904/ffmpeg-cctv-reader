package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prakasa1904/ffmpeg-cctv-reader/command"
)

func runningFFMPEG(ctx context.Context, ffmpegCmd string) error {
	command := command.NewCommand(
		command.WithContext(ctx),
		command.WithName("ffmpeg"),
		command.WithCommand(ffmpegCmd),
	).SetStdout(os.Stdout).SetStderr(os.Stderr).SetEnv(os.Environ())

	return command.Run()
}
func runningMediaMTX(ctx context.Context, mediaMTXCmd string) error {
	command := command.NewCommand(
		command.WithContext(ctx),
		command.WithName("./publisher/mediamtx"),
		command.WithCommand(mediaMTXCmd),
	).SetStdout(os.Stdout).SetStderr(os.Stderr).SetEnv(os.Environ())

	return command.Run()
}

func main() {
	// Create a context with a timeout of 5 seconds.
	ctxFFMPEG, cancelFMPEG := context.WithCancel(context.Background())
	ctxMediaMTX, cancelMediaMTX := context.WithCancel(context.Background())

	cctvCamera := os.Getenv("CCTV_CAMERA")
	rtspServer := os.Getenv("RTSP_SERVER")

	if cctvCamera == "" || rtspServer == "" {
		fmt.Println("CCTV_CAMERA and RTSP_SERVER environment variables must be set.")
		return
	}
	ffmpeArgs := fmt.Sprintf("-i %s -c:v copy -c:a copy -f rtsp %s", cctvCamera, rtspServer)
	mediaMTXArgs := "publisher/mediamtx.yml"

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Println("Starting mediamtx command")
	go func() {
		err := runningMediaMTX(ctxMediaMTX, mediaMTXArgs)
		if err != nil {
			log.Printf("Error running mediamtx: %v\n", err)
		} else {
			log.Println("MediaMTX command executed successfully.")
		}
	}()

	time.Sleep(5 * time.Second)
	log.Println("Starting ffmpeg command")
	go func() {
		err := runningFFMPEG(ctxFFMPEG, ffmpeArgs)
		if err != nil {
			log.Printf("Error running ffmpeg: %v\n", err)
		} else {
			log.Println("FFmpeg command executed successfully.")
		}
	}()

	select {
	case <-done:
		fmt.Println("Received interrupt signal. Stopping command...")
		// Cancel the context to signal the command to stop.
		cancelFMPEG()
		cancelMediaMTX()
		ctxFFMPEG.Done()
		ctxMediaMTX.Done()
	case err := <-done:
		if err != nil {
			log.Printf("Error running ffmpeg and mediamtx: %v\n", err)
		} else {
			log.Println("FFmpeg and MediaMTX command executed successfully.")
		}
	}
}
