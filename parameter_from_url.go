package gohttplib

import (
	"net/http"
	"strconv"
)

func IntParameterFromURLInRequest(r *http.Request, name string) *int {
	s := GetParameterFromURLInRequest(r, name)
	if s == nil {
		return nil
	}
	i64, err := strconv.ParseInt(*s, 10, 64)
	if err != nil {
		return nil
	}
	i := int(i64)
	return &i
}

func Int64ParameterFromURLInRequest(r *http.Request, name string) *int64 {
	s := GetParameterFromURLInRequest(r, name)
	if s == nil {
		return nil
	}
	i64, err := strconv.ParseInt(*s, 10, 64)
	if err != nil {
		return nil
	}
	return &i64
}

func BoolParameterFromURLInRequest(r *http.Request, name string) *bool {
	s := GetParameterFromURLInRequest(r, name)
	if s == nil {
		return nil
	}
	b, err := strconv.ParseBool(*s)
	if err != nil {
		return nil
	}
	return &b
}

func FloatParameterFromURLInRequest(r *http.Request, name string) *float64 {
	s := GetParameterFromURLInRequest(r, name)
	if s == nil {
		return nil
	}
	float, err := strconv.ParseFloat(*s, 64);
	if err != nil {
		return nil
	}
	return &float
}
