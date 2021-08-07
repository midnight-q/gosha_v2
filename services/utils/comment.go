package utils

import "strings"

func ParseComment(list []string) string {
	return strings.Replace(strings.Join(list, "\n"), "//", "", -1)
}
