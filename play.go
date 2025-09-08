package main

import (
	"os/exec"
	"fmt"
	"os"
)


func playVid(videoId string) {
	url := "https://www.youtube.com/watch?v=" + videoId


	cmd := exec.Command("mpv", "--no-resume-playback", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	
	cmd.Env = os.Environ()
	if err := cmd.Run(); err != nil {
			fmt.Println("Error playing video:", err)
	}
}
