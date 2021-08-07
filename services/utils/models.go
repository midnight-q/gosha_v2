package utils

import (
	"strings"
)

func CheckIsFilter(name string) bool {
	return strings.Contains(name, "Filter")
}

func IsIdField(filedName string) bool {
	if strings.ToLower(filedName) == "id" {
		return true
	}
	return false
}
