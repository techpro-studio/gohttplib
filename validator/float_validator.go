package validator

import (
	"github.com/techpro-studio/gohttplib"
)

type FloatRange struct {
	Upper  *float64
	Bottom *float64
}

func FloatValidator(key string) Validator {
	return func(value interface{}) error {
		_, ok := value.(float64)
		if !ok {
			return gohttplib.NewError(key, "Should be float", "TYPE_ERROR", []string{"float"})
		}
		return nil
	}
}

func FloatInRangeValidator(key string, floatRange FloatRange) Validator {
	return func(value interface{}) error {
		float := value.(float64)
		err := gohttplib.NewError(key, "Invalid float", "FLOAT_RANGE_ERROR", nil)
		if floatRange.Upper != nil && *floatRange.Upper < float {
			return err
		}
		if floatRange.Bottom != nil && *floatRange.Bottom > float {
			return err
		}
		return nil
	}
}

func RequiredFloatValidators(key string, validators ...Validator) []Validator {
	arr := []Validator{NotEmptyValidator(key), FloatValidator(key)}
	return append(arr, validators...)
}
