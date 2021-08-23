package common

import (
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func ParseInt(req *http.Request, name string) int {
	num, _ := strconv.Atoi(req.FormValue(name))
	return num
}
func ParseFloat(req *http.Request, name string) float64 {
	num, _ := strconv.ParseFloat(req.FormValue(name), 64)
	return num
}
func ParseString(req *http.Request, name string) string {
	return req.FormValue(name)
}
func ParseBool(req *http.Request, name string) bool {
	b, _ := strconv.ParseBool(req.FormValue(name))
	return b
}
func ParseUuid(req *http.Request, name string) uuid.UUID {
	num, _ := uuid.Parse(req.FormValue(name))
	return num
}
func ParseTime(req *http.Request, name string) time.Time {
	// ParseTime Expected format "2006-01-02T15:04:05Z07:00"
	num, _ := time.Parse(time.RFC3339, req.FormValue(name))
	return num
}

func ParseIntPointer(req *http.Request, name string) *int {
	s := req.FormValue(name)
	if len(s) < 1 {
		return nil
	}
	num, _ := strconv.Atoi(s)
	return &num
}
func ParseFloatPointer(req *http.Request, name string) *float64 {
	s := req.FormValue(name)
	if len(s) < 1 {
		return nil
	}
	num, _ := strconv.ParseFloat(s, 64)
	return &num
}
func ParseBoolPointer(req *http.Request, name string) *bool {
	s := req.FormValue(name)
	if len(s) < 1 {
		return nil
	}
	b, _ := strconv.ParseBool(s)
	return &b
}
func ParseStringPointer(req *http.Request, name string) *string {
	s := req.FormValue(name)
	if len(s) < 1 {
		return nil
	}
	return &s
}
func ParseUuidPointer(req *http.Request, name string) *uuid.UUID {
	s := req.FormValue(name)
	if len(s) < 1 {
		return nil
	}
	num, _ := uuid.Parse(s)
	return &num
}
func ParseTimePointer(req *http.Request, name string) *time.Time {
	s := req.FormValue(name)
	if len(s) < 1 {
		return nil
	}
	num, _ := time.Parse(time.RFC3339, s)
	return &num
}

func ParseStringArray(req *http.Request, name string) (res []string) {
	if req.Form == nil {
		req.ParseForm()
	}
	for _, s := range req.Form[name] {
		res = append(res, s)
	}
	return
}
func ParseIntArray(req *http.Request, name string) (res []int) {
	strArray := ParseStringArray(req, name)
	for _, s := range strArray {
		num, _ := strconv.Atoi(s)
		res = append(res, num)
	}
	return
}
func ParseFloatArray(req *http.Request, name string) (res []float64) {
	strArray := ParseStringArray(req, name)
	for _, s := range strArray {
		num, _ := strconv.ParseFloat(s, 64)
		res = append(res, num)
	}
	return
}
func ParseBoolArray(req *http.Request, name string) (res []bool) {
	strArray := ParseStringArray(req, name)
	for _, s := range strArray {
		num, _ := strconv.ParseBool(s)
		res = append(res, num)
	}
	return
}
func ParseUuidArray(req *http.Request, name string) (res []uuid.UUID) {
	strArray := ParseStringArray(req, name)
	for _, s := range strArray {
		num, _ := uuid.Parse(s)
		res = append(res, num)
	}
	return
}
func ParseTimeArray(req *http.Request, name string) (res []time.Time) {
	strArray := ParseStringArray(req, name)
	for _, s := range strArray {
		num, _ := time.Parse(time.RFC3339, s)
		res = append(res, num)
	}
	return
}
