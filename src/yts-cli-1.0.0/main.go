package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

func main() {

	tmpScript, err := writeTempScript()
	if err != nil {
		fmt.Println("Failed to create temp script: ", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	var styleQuery = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#8FA31E"))

	var maxResStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#556B2F"))

	var headerStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#c6d870"))

	fmt.Print(styleQuery.Render("Search for a video title: "))
	query, _ := reader.ReadString('\n')
	query = strings.TrimSpace(query)

	fmt.Print(maxResStyle.Render("Number of Results: "))
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

	var resultsStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#5d688a")).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")).
		MaxWidth(117).
		Align(lipgloss.Left)

	var videoTitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#556b2f"))

	var URLStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#8fa31e"))

	for i := range idList {
		if i < len(titleList) {
			title := videoTitleStyle.Render(fmt.Sprintf("Video Title: %s", titleList[i]))
			url := URLStyle.Render(fmt.Sprintf("URL: https://www.youtube.com/watch?v=%s", idList[i]))

			block := resultsStyle.Render(title + "\n" + url)

			fmt.Println(block)
			fmt.Println()
		}
	}
	
	defer os.Remove("videoIds.txt")
	defer os.Remove("videoTitles.txt")
	defer os.Remove("yInitialData.json")
	defer os.Remove("combined.txt")
	defer os.Remove(tmpScript)
}
