# YTS-CLI

![[Logo]](yts-cli-logo.png)

YTS-CLI is a CLI YouTube Web Scraper, a lightweight command line tool to search YouTube and get video Titles and URLs WITHOUT using the YouTube API.

## Features:
- Search YouTube directly from terminal.
- Specify the number of results you want to see.
- Outputs video titles and URLs in neat boxes.

## Installation: 

From **AUR** (Arch Linux User Repository)

```bash
yay -S yts-cli
```

If you want to build it from source:

```bash
git clone https://github.com/IxBlazarxI/yts-cli.git
cd yts-cli
makepkg -si
```

## Usage:

Run the CLI:

```bash
yts-cli
```

Example:

![[yts-cli in action]](demo.png)

## Credits:

Credits to charmbracelet/lipgloss library for the box styling. (https://github.com/charmbracelet/lipgloss)
