package validator

import (
	"github.com/wolvesstudio/gohttplib"
)

type IntRange struct {
	Upper  *int
	Bottom *int
}

type Int64Range struct {
	Upper  *int64
	Bottom *int64
}

func IntValidator(key string) Validator {
	return func(value interface{}) error {
		_, ok := value.(int)
		_, ok = value.(int64)
		if !ok {
			return gohttplib.NewError(key, " Should be int", "TYPE_ERROR", []string{"int"})
		}
		return nil
	}
}


func RequiredIntValidators(key string, validators ...Validator) []Validator {
	arr := []Validator{NotEmptyValidator(key), IntValidator(key)}
	return append(arr, validators...)
}

func IntInRangeValidator(key string, intRange IntRange) Validator {
	return func(value interface{}) error {
		intValue := value.(int)
		err := gohttplib.NewError(key, "Invalid int", "INT_RANGE_ERROR", nil)
		if intRange.Upper != nil && *intRange.Upper < intValue {
			return err
		}
		if intRange.Bottom != nil && *intRange.Bottom > intValue {
			return err
		}
		return nil
	}
}


func Int64InRangeValidator(key string, intRange Int64Range) Validator {
	return func(value interface{}) error {
		intValue := value.(int64)
		err := gohttplib.NewError(key, "Invalid int64", "INT64_RANGE_ERROR", nil)
		if intRange.Upper != nil && *intRange.Upper < intValue {
			return err
		}
		if intRange.Bottom != nil && *intRange.Bottom > intValue {
			return err
		}
		return nil
	}
}

