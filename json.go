package gohttplib

import (
	"bufio"
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, v any, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	bw := bufio.NewWriter(w)
	enc := json.NewEncoder(bw)
	enc.SetEscapeHTML(false)

	if err := enc.Encode(v); err != nil {
		return // log if needed
	}

	bw.Flush()
}

func WriteJsonOrError(w http.ResponseWriter, value interface{}, code int, err error) {
	if err != nil {
		SafeConvertToServerError(err).Write(w)
		return
	}
	WriteJson(w, value, code)
}
