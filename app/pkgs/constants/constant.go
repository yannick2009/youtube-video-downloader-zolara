package constants

// yt-dlp commands
const (
	YTDL_COMMAND        = "yt-dlp"
	YTDL_FORMAT_FLAG    = "-f"
	YTDL_COMMAND_FORMAT = "bestvideo[ext=mp4]+bestaudio[ext=m4a]/best[ext=mp4]"
	YTDL_OUTPUT_FLAG    = "-o"
	YTDL_OUTPUT_FORMAT  = "videos/%(title)s.%(ext)s"
)

// file paths
const (
	FILES_FOLDER_PATH         = "assets/"
	COMPARING_TOPIC_FILE_PATH = FILES_FOLDER_PATH + "comparing-object-or-people.txt"
)
