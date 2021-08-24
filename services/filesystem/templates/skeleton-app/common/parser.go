package common

import (
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func ParseIntFromRequest(req *http.Request, name string) int {
	num, _ := strconv.Atoi(req.FormValue(name))
	return num
}
func ParseFloatFromRequest(req *http.Request, name string) float64 {
	num, _ := strconv.ParseFloat(req.FormValue(name), 64)
	return num
}
func ParseBoolFromRequest(req *http.Request, name string) bool {
	b, _ := strconv.ParseBool(req.FormValue(name))
	return b
}
func ParseStringFromRequest(req *http.Request, name string) string {
	return req.FormValue(name)
}
func ParseUuidFromRequest(req *http.Request, name string) uuid.UUID {
	num, _ := uuid.Parse(req.FormValue(name))
	return num
}
func ParseTimeFromRequest(req *http.Request, name string) time.Time {
	// ParseTime Expected format "2006-01-02T15:04:05Z07:00"
	num, _ := time.Parse(time.RFC3339, req.FormValue(name))
	return num
}

func ParseIntPointerFromRequest(req *http.Request, name string) *int {
	s := req.FormValue(name)
	if len(s) < 1 {
		return nil
	}
	num, _ := strconv.Atoi(s)
	return &num
}
func ParseFloatPointerFromRequest(req *http.Request, name string) *float64 {
	s := req.FormValue(name)
	if len(s) < 1 {
		return nil
	}
	num, _ := strconv.ParseFloat(s, 64)
	return &num
}
func ParseBoolPointerFromRequest(req *http.Request, name string) *bool {
	s := req.FormValue(name)
	if len(s) < 1 {
		return nil
	}
	b, _ := strconv.ParseBool(s)
	return &b
}
func ParseStringPointerFromRequest(req *http.Request, name string) *string {
	s := req.FormValue(name)
	if len(s) < 1 {
		return nil
	}
	return &s
}
func ParseUuidPointerFromRequest(req *http.Request, name string) *uuid.UUID {
	s := req.FormValue(name)
	if len(s) < 1 {
		return nil
	}
	num, _ := uuid.Parse(s)
	return &num
}
func ParseTimePointerFromRequest(req *http.Request, name string) *time.Time {
	s := req.FormValue(name)
	if len(s) < 1 {
		return nil
	}
	num, _ := time.Parse(time.RFC3339, s)
	return &num
}

func ParseStringArrayFromRequest(req *http.Request, name string) (res []string) {
	if req.Form == nil {
		req.ParseForm()
	}
	for _, s := range req.Form[name] {
		res = append(res, s)
	}
	return
}
func ParseIntArrayFromRequest(req *http.Request, name string) (res []int) {
	strArray := ParseStringArrayFromRequest(req, name)
	for _, s := range strArray {
		num, _ := strconv.Atoi(s)
		res = append(res, num)
	}
	return
}
func ParseFloatArrayFromRequest(req *http.Request, name string) (res []float64) {
	strArray := ParseStringArrayFromRequest(req, name)
	for _, s := range strArray {
		num, _ := strconv.ParseFloat(s, 64)
		res = append(res, num)
	}
	return
}
func ParseBoolArrayFromRequest(req *http.Request, name string) (res []bool) {
	strArray := ParseStringArrayFromRequest(req, name)
	for _, s := range strArray {
		num, _ := strconv.ParseBool(s)
		res = append(res, num)
	}
	return
}
func ParseUuidArrayFromRequest(req *http.Request, name string) (res []uuid.UUID) {
	strArray := ParseStringArrayFromRequest(req, name)
	for _, s := range strArray {
		num, _ := uuid.Parse(s)
		res = append(res, num)
	}
	return
}
func ParseTimeArrayFromRequest(req *http.Request, name string) (res []time.Time) {
	strArray := ParseStringArrayFromRequest(req, name)
	for _, s := range strArray {
		num, _ := time.Parse(time.RFC3339, s)
		res = append(res, num)
	}
	return
}
