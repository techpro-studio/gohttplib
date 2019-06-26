package gohttplib

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const paramsKey = "parameters"

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		params := make(map[string]string)
		for _, value := range ps {
			params[value.Key] = value.Value
		}
		h.ServeHTTP(w, SetInContext(params, paramsKey, r))
	}
}

func SetInContext(value interface{}, key interface{}, req *http.Request) *http.Request {
	ctx := context.WithValue(req.Context(), key, value)
	return req.WithContext(ctx)
}

func GetParameterFromURLInRequest(r *http.Request, key string) *string {
	params := r.Context().Value(paramsKey).(map[string]string)
	var value string
	if len(params) > 0 {
		value = params[key]
	}
	if len(value) == 0 {
		value = r.URL.Query().Get(key)
	}

	if len(value) == 0 {
		return nil
	}
	return &value
}
