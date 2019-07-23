package gohttplib

import (
	"encoding/json"
	"net/http"
)

func GetBody(req *http.Request) (map[string]interface{}, error) {
	decoder := json.NewDecoder(req.Body)
	var _map map[string]interface{}
	err := decoder.Decode(&_map)
	defer req.Body.Close()
	if err != nil {
		return nil, HTTP400(err.Error())
	}
	return _map, nil
}
