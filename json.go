package gohttplib

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, value interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	bytes, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(bytes)
	if err != nil {
		panic(err)
	}
}

func WriteJsonOrError(w http.ResponseWriter, value interface{}, code int, err error){
	if err != nil {
		SafeConvertToServerError(err).Write(w)
		return
	}
	WriteJson(w, value, code)
}