package utils

import "strings"

func GetNameForNewApp(path string) string {
	path = strings.Trim(path, "/")
	path = strings.Trim(path, "\\")
	arr := strings.Split(path, "/")
	if len(arr) == 1 {
		arr = strings.Split(path, "\\")
	}
	return arr[len(arr)-1]
}
