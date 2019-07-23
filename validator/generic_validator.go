package validator

import (
	"github.com/techpro-studio/gohttplib"
)

type Validator func(value interface{}) error

func NotEmptyValidator(key string) Validator {
	return func(value interface{}) error {
		if value == nil {
			return gohttplib.NewError(key, "Field is required", "REQUIRED_FIELD_ERROR", nil)
		}
		return nil
	}
}

func BoolValidator(key string) Validator {
	return func(value interface{}) error {
		_, ok := value.(bool)
		if !ok {
			return gohttplib.NewError(key, "Should be bool", "TYPE_ERROR", []string{"bool"})
		}
		return nil
	}
}

func ArrayValidator(key string) Validator {
	return func(value interface{}) error {
		_, ok := value.([]interface{})
		if !ok {
			return gohttplib.NewError(key, "Should be array", "TYPE_ERROR", []string{"array"})
		}
		return nil
	}
}

func RequiredBoolValidators(key string, validators ...Validator) []Validator {
	arr := []Validator{NotEmptyValidator(key), BoolValidator(key)}
	return append(arr, validators...)
}

func ValidateValue(value interface{}, validators []Validator) []gohttplib.Error {
	errs := []gohttplib.Error{}
	for _, validator := range validators {
		err := validator(value)
		if err != nil {
			errs = append(errs, err.(gohttplib.Error))
			break
		}
	}
	return errs
}

type VMap map[string][]Validator

func ValidateMap(dictionary map[string]interface{}, validatorMap VMap) []gohttplib.Error {
	errs := []gohttplib.Error{}
	for key, validators := range validatorMap {
		errs = append(errs, ValidateValue(dictionary[key], validators)...)
	}
	return errs
}
