package internal

import (
	"embed"
)

//go:embed assets/*.txt
var FilesFolder embed.FS
 