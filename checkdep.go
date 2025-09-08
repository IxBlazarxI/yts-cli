package main

import (
	"os"
	"github.com/charmbracelet/lipgloss"
	"os/exec"
	"fmt"
)

var (
	 checkDepStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffd700"))
	 installDepStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff87"))
	 depFoundStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#6ad6ff"))
)

func checkDep() {
	fmt.Println(checkDepStyle.Render("Checking dependencies..."))
	deps := []string{"mpv", "yt-dlp"}
	var missing []string

	for _, cmd := range deps {
		if _, err := exec.LookPath(cmd); err != nil {
			missing = append(missing, cmd)
		}
	}

	if len(missing) > 0 {
		fmt.Println(installDepStyle.Render("Missing dependencies:"))
		for _, cmd := range missing {
			fmt.Printf(installDepStyle.Render("Run: sudo pacman -S %s\n", cmd))
		}
		os.Exit(1)
	}
}
