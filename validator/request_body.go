package validator

import (
	"net/http"

	"github.com/techpro-studio/gohttplib"
)

func ValidateBody(body map[string]interface{}, validatorMap VMap) (map[string]interface{}, error) {
	errs := ValidateMap(body, validatorMap)
	if len(errs) > 0 {
		return nil, &gohttplib.ServerError{400, gohttplib.Errors{Errors: errs}}
	}
	return body, nil
}

func GetValidatedBody(req *http.Request, validatorMap VMap) (map[string]interface{}, error) {
	body, err := gohttplib.GetBody(req)
	if err != nil {
		return nil, err
	}
	return ValidateBody(body, validatorMap)
}
