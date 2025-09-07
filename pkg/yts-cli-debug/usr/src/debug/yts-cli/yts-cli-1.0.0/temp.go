package main

import (
	"os"
)

const bashScript = `#!/bin/bash

query="$1"
max_results="$2"

search_query=$(echo "$query" | sed 's/ /+/g')

# Download the page and store in yInitialData.json
curl -s "https://www.youtube.com/results?search_query=$search_query" -o yInitialData.json

# Extract videoId and title with this insane regex
grep -oP '"videoRenderer":\s*\{[^}]*"videoId":"[^"]+".*?"title":\{"runs":\[\{"text":"[^"]+' yInitialData.json \
  | head -n "$max_results" > combined.txt

# Splitting IDs and Video Titles into different files
awk -F'"videoId":"' '{print $2}' combined.txt | awk -F'"' '{print $1}' > videoIds.txt
awk -F'"text":"' '{print $2}' combined.txt | awk -F'"' '{print $1}' > videoTitles.txt
`

func writeTempScript() (string, error) {
	tmpPath := os.TempDir() + "/ytscraper_temp.sh"
	if err := os.WriteFile(tmpPath, []byte(bashScript), 0755); err != nil {
		return "", err
	}
	return tmpPath, nil
}
