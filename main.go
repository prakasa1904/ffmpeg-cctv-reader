package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/prakasa1904/ffmpeg-cctv-reader/command"
)

func runningMediaMTX(ctx context.Context, mediaMTXCmd string) error {
	command := command.NewCommand(
		command.WithContext(ctx),
		command.WithName("./publisher/mediamtx"),
		command.WithCommand(mediaMTXCmd),
	).SetStdout(os.Stdout).SetStderr(os.Stderr).SetEnv(os.Environ())

	return command.Run()
}

func main() {
	ctxMediaMTX, cancelMediaMTX := context.WithCancel(context.Background())
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

	select {
	case <-done:
		fmt.Println("Received interrupt signal. Stopping command...")
		// Cancel the context to signal the command to stop.
		cancelMediaMTX()
		ctxMediaMTX.Done()
	case err := <-done:
		if err != nil {
			log.Printf("Error running mediamtx: %v\n", err)
		} else {
			log.Println("MediaMTX command executed successfully.")
		}
	}
}
