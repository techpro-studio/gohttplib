package validator

import (
	"fmt"
	"github.com/johngb/langreg"
	"github.com/wolvesstudio/gohttplib"
	"net/url"
	"strings"
)

func StringValidator(key string) Validator {
	return func(value interface{}) error {
		_, ok := value.(string)
		if !ok {
			return gohttplib.NewError(key, "Should be string", "TYPE_ERROR", []string{"string"})
		}
		return nil
	}
}

func StringLengthValidator(length int, key string) Validator {

	return func(value interface{}) error {
		stringValue := value.(string)
		if len(stringValue) < length {
			return gohttplib.NewError(key, fmt.Sprintf("%@ should be minimum %d characters", strings.ToUpper(key), length),
				"STRING_LENGTH_ERROR", []string{key, fmt.Sprintf("%d", length)})

		}
		return nil
	}
}

func StringArrayValidator(key string, each []Validator) Validator {
	return func(value interface{}) error {
		values := value.([]interface{})
		strArr := []string{}
		for _, item := range values {
			str, ok := item.(string)
			if !ok {
				return gohttplib.NewError(key, "Should be string in array", "TYPE_ERROR", []string{"string", "array"})
			}
			strArr = append(strArr, str)
		}
		for _, item := range strArr {
			for _, validator := range each {
				err := validator(item)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func LanguageValidator(key string) Validator {
	return func(value interface{}) error {
		stringValue := value.(string)
		if !langreg.IsValidLanguageCode(stringValue) {
			return gohttplib.NewError(key, "Invalid language", "INVALID_LANGUAGE_ERROR", []string{stringValue})

		}
		return nil
	}
}

func URLValidator(key string) Validator {
	return func(value interface{}) error {
		stringValue := value.(string)

		_, err := url.Parse(stringValue)
		if err != nil {
			return gohttplib.NewError(key, "Invalid url", "INVALID_URL_ERROR", nil)
		}
		return nil
	}
}

func SexValidator(key string) Validator {
	return StringContainsValidator(key, []string{"male", "female"})
}

func StringContainsValidator(key string, values []string) Validator {
	return func(value interface{}) error {
		stringValue := value.(string)
		contains := false
		for _, item := range values {
			if item == stringValue {
				contains = true
				break
			}
		}
		if !contains {
			return gohttplib.NewError(key, fmt.Sprintf("Invalid %s", key),
				fmt.Sprintf("INVALID_%s_ERROR", strings.ToUpper(key)), nil)
		}
		return nil
	}
}



func CountryValidator(key string) Validator {
	return func(value interface{}) error {
		stringValue := value.(string)
		if !langreg.IsValidRegionCode(stringValue) {
			return gohttplib.NewError(key, "Invalid country", "INVALID_COUNTRY_ERROR", nil)
		}
		return nil
	}
}

func RequiredStringValidators(key string, validators ...Validator) []Validator {
	arr := []Validator{NotEmptyValidator(key), StringValidator(key)}
	return append(arr, validators...)
}
