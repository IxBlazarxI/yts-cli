#!/bin/bash

query="$1"
max_results="$2"

# Replace spaces with + for YouTube search
search_query=$(echo "$query" | sed 's/ /+/g')

# Download the page
curl -s "https://www.youtube.com/results?search_query=$search_query" -o yInitialData.json

# Extract videoId and title (runs style)
grep -oP '"videoRenderer":\s*\{[^}]*"videoId":"[^"]+".*?"title":\{"runs":\[\{"text":"[^"]+' yInitialData.json \
  | head -n "$max_results" > combined.txt

# Split IDs
awk -F'"videoId":"' '{print $2}' combined.txt | awk -F'"' '{print $1}' > videoIds.txt

# Split Titles
awk -F'"text":"' '{print $2}' combined.txt | awk -F'"' '{print $1}' > videoTitles.txt
