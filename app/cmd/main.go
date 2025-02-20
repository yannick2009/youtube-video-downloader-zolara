package main

import (
	"yt-video-downloader/internal"
	"yt-video-downloader/pkgs/constants"
)

func main() {
	// read the file in the directory assets (embedded)
	//read the content of the file
	// clean each youtube short lint to remove /short/ and get the video id
	// for each link run the youtube-dl command to download the video
	// save the video in the directory videos/<topic>/<video_id>.mp4
	fileByte, err := internal.FilesFolder.ReadFile(constants.COMPARING_TOPIC_FILE_PATH)
	if err != nil {
		panic(err)
	}
	file := string(fileByte)
	linksArr := internal.CleanYTLinks(file)
	internal.DownloadVideos(linksArr)
}
