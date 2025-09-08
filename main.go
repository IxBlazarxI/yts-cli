package main

import (
	"fmt"
	"os/signal"
	"bufio"
	"syscall"
	"os"
	"os/exec"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	promptStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00ff87"))
	
	resultsStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#a3a0f0")).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#6f6dc7")).
		MaxWidth(117).
		Align(lipgloss.Left)

	 videoTitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#ffd580"))

	 URLStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#80dfff"))

	headerStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#ff9f80"))

	videoChooseStyle = lipgloss.NewStyle().
		Bold(true). 
		Foreground(lipgloss.Color("#ffcf6f"))
)

func main() {

	checkDep()
	tmpScript, err := writeTempScript()
	if err != nil {
		fmt.Println("Failed to create temp script: ", err)
		return
	}

	tempFiles := []string {
		"videoIds.txt",
		"videoTitles.txt",
		"yInitialData.json",
		"combined.txt",
		tmpScript,
	}

	defer cleanUp(tempFiles...)

	// Remove tmp files if CTRL + C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func () {
		<-c
		cleanUp(tempFiles...)
		os.Exit(1)
	}()

	reader := bufio.NewReader(os.Stdin)


	fmt.Print(promptStyle.Render("Search for a video title: "))
	query, _ := reader.ReadString('\n')
	query = strings.TrimSpace(query)

	fmt.Print(promptStyle.Render("Number of Results: "))
	var maxResults string
	fmt.Scan(&maxResults)
	fmt.Print("\n")


	cmd := exec.Command(tmpScript, query, maxResults)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running scraper: ", err)
		return 
	}

	ids, _ := os.ReadFile("videoIds.txt")
	titles, _ := os.ReadFile("videoTitles.txt")

	idList := strings.Split(strings.TrimSpace(string(ids)), "\n")
	titleList := strings.Split(strings.TrimSpace(string(titles)), "\n")

	fmt.Println(headerStyle.Render("Search Results:\n"))


	for i := range idList {
		if i < len(titleList) {
			title := videoTitleStyle.Render(fmt.Sprintf("%d. Video Title: %s", i+1, titleList[i]))
			url := URLStyle.Render(fmt.Sprintf("URL: https://www.youtube.com/watch?v=%s", idList[i]))

			block := resultsStyle.Render(title + "\n" + url)

			fmt.Println(block)
			fmt.Println()
		}
	}

	var choice int

	fmt.Print(videoChooseStyle.Render("Choose video to play: "))
	fmt.Scan(&choice)

	if choice > len(idList) || choice < 1 {
		fmt.Println("Please select from available videos.")
	}

	selectedId := idList[choice - 1]
	playVid(selectedId)

}

func cleanUp(files ...string) {
	for _, f := range files {
		os.Remove(f)
	}
}
