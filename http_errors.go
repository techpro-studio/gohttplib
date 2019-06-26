package gohttplib

import (
	"fmt"
)

func HTTP400(description *string) *ServerError {
	return NewServerError(400, "INVALID_REQUEST", description, nil, nil)
}

func HTTP401() *ServerError {
	return NewServerError(401, "UNAUTHORIZED", nil , nil, nil)
}

func HTTP403(code string) *ServerError {
	return NewServerError(403, code, nil , nil , nil)
}

func DefaultHTTP403() *ServerError {
	return  HTTP403("PERMISSION_DENIED")
}

func HTTP404(id string) *ServerError {
	return NewServerError(404, "NOT_FOUND", nil, nil, []string{id})
}

func HTTP500(err interface{}) *ServerError {
	return NewServerError(500, "FATAL", nil, nil, []string{fmt.Sprintf("%v", err)})
}


