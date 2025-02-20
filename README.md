# YouTube Video Downloader

A Go tool to automatically download YouTube Shorts videos.
Exclusively for the ZOLARA project (ESP Epitech2025)

## Prerequisites

- Go 1.24 or higher
- yt-dlp
- Nix (optional, for development environment)

## Installation

1. Clone the repository:
```bash
git clone <repo-url>
cd app
```

2. Installation with Nix (recommended):
```bash
nix develop
```

Or manually install dependencies:
- Install Go 1.24
- Install yt-dlp

## Usage

1. Add your YouTube Shorts links in the `internal/assets/comparing-object-or-people.txt` file
2. Run the program:
```bash
go run cmd/main.go
```

Videos will be downloaded to the `videos/` folder in MP4 format.

## Project Structure

```
app/
├── cmd/          # Application entry point
├── internal/     # Internal code
├── pkgs/         # Reusable packages
└── videos/       # Video output folder
```

## Features

- Automatic YouTube Shorts downloading
- High-quality MP4 output format
- Link management from text file
- Uses yt-dlp for optimal compatibility

## License

MIT License

