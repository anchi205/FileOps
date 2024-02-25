package utils

import "strings"

func TrimFileName(filename string) string {
	stringSlice := strings.Split(filename, "/")
	filename = stringSlice[len(stringSlice)-1]
	return filename
}
