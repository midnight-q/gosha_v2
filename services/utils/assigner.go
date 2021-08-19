package utils

import "strings"

func IsTypeFromDb(name string) bool {
	return strings.HasSuffix(name, "TypeFromDb")
}

func IsmDbFromType(name string) bool {
	return strings.HasSuffix(name, "DbFromType")
}
