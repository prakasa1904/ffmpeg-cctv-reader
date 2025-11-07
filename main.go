package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	user := os.Getenv("USERNAME")
	pass := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if user == "" || pass == "" {
		fmt.Println("USERNAME and PASSWORD environment variables must be set.")
		return
	}

	ffmpegCmd := fmt.Sprintf("ffmpeg -i rtsp://%s:%s@%s:%s/live/ch00_0 -c:v copy -c:a copy -hls_time 2 -hls_list_size 5 -hls_flags delete_segments -start_number 1 ./stream/output.m3u8", user, pass, host, port)
	args := strings.Fields(ffmpegCmd)

	cmd := exec.Command(args[0], args[1:]...)

	// Connect the command's stdout to the current process's stdout
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr // Optionally, connect stderr as well

	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running ffmpeg: %v\n", err)
	} else {
		fmt.Println("FFmpeg command executed successfully.")
	}
}
