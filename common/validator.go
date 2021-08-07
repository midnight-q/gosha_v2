package common

import (
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"unicode"
)

func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

func ValidatePassword(password string) bool {
	count := 0
	hasChar := false
	hasNum := false

	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			count++
			hasNum = true
		case unicode.IsLetter(c):
			count++
			hasChar = true
		default:
			return false
		}
	}

	return hasChar && hasNum && count > 7
}

func ValidateApplicationName(name string) bool {
	// The name may contain lowercase letters ('a' through 'z'), numbers, and underscores ('_'). Name parts may only start with letters.
	Re := regexp.MustCompile(`^([a-z]{1,})+([a-z\d-_]*)[a-z\d]$`)
	return Re.MatchString(name)
}

func ValidateMobile(phone string) bool {
	Re := regexp.MustCompile(`^[+][0-9]{11,}`)
	return Re.MatchString(phone)
}

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func CheckInArray(item interface{}, array interface{}) bool {
	exist, _ := InArray(item, array)
	return exist
}

func getFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}

// MyCaller returns the caller of the function that called it :)
func MyCaller() string {
	// Skip GetCallerFunctionName and the function to get the caller of
	return getFrame(2).Function
}

func UniqueStringArray(slice []string) []string {
	// create a map with all the values as key
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}

func UniqueIntArray(slice []int) []int {
	// create a map with all the values as key
	uniqMap := make(map[int]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]int, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}

func ReverseBytes(in []byte) []byte {
	res := []byte{}

	for i := len(in) - 1; i >= 0; i-- {
		res = append(res, in[i])
	}
	return res
}

func IntArrayToStringArray(in []int) (res []string) {
	for _, val := range in {
		res = append(res, strconv.Itoa(val))
	}
	return
}

func StringArrayToIntArray(in []string) (res []int) {
	for _, val := range in {
		if v, err := strconv.Atoi(val); err == nil {
			res = append(res, v)
		}
	}
	return
}
