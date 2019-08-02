package gohttplib

import (
	"fmt"
	"net/http"
)

type Errors struct {
	Errors []Error `json:"errors"`
}

type ServerError struct {
	StatusCode int
	Errors     Errors
}

func (err ServerError) Error() string {
	return err.Errors.Error()
}

func SafeConvertToServerError(err error)*ServerError{
	serverError, ok := err.(ServerError)
	if ok  {
		return &serverError
	}
	return NewServerError(400, "UNDEFINED", err.Error(), "", nil)
}

func (err ServerError) Write(w http.ResponseWriter) {
	WriteJson(w, err.Errors, err.StatusCode)
}

type Error struct {
	Key         string   `json:"key, omitempty"`
	Description string   `json:"description, omitempty"`
	Code        string   `json:"code"`
	Args        []string `json:"args, omitempty"`
}

func NewError(key string, description string, code string, args []string) *Error {
	return &Error{Key: key, Description: description, Code: code, Args: args}
}

func (err Error) WriteWithCode(code int, w http.ResponseWriter) {
	ServerError{code, Errors{[]Error{err}}}.Write(w)
}

func (err Error) AsServerError(code int) error {
	return ServerError{code, Errors{[]Error{err}}}
}

func NewServerError(statusCode int, code string, description string, key string, args []string) *ServerError {
	return &ServerError{
		StatusCode: statusCode,
		Errors: Errors{Errors: []Error{{
			Key:         key,
			Description: description,
			Code:        code,
			Args:        args,
		}}}}
}

func (err Error) Error() string {
	return err.Code
}

func (errs Errors) Error() string {
	return fmt.Sprintf("Occured %d errors", len(errs.Errors))
}
