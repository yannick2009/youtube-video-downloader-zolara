package internal

import "strings"

const SHORTS_URL_PATTERN = "/shorts/"

func CleanYTLinks(s string) []string {
	var arr []string
	for line := range strings.Lines(s) {
		arr = append(arr, line)
	}
	return arr
}
