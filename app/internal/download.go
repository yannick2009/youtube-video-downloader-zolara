package internal

import (
	"log"
	"os/exec"
	"yt-video-downloader/pkgs/constants"
)

func DownloadVideos(links []string) {
	log.Printf("\033[1;34m%s\033[0m", "Downloading videos...")

	cmd := exec.Command(constants.YTDL_COMMAND)
	args := []string{constants.YTDL_FORMAT_FLAG, constants.YTDL_COMMAND_FORMAT, constants.YTDL_OUTPUT_FLAG, constants.YTDL_OUTPUT_FORMAT}
	args = append(args, links...)
	cmd.Args = append(cmd.Args, args...)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	log.Printf("\033[1;32m%s\033[0m", "Videos downloaded successfully!")
}
