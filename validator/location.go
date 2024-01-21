package validator

func LongitudeValidators(key string) []Validator {
	upper := 180.0
	bottom := -180.0
	return RequiredFloatValidators(key, FloatInRangeValidator(key, FloatRange{&upper, &bottom}))
}

func LatitudeValidators(key string) []Validator {
	upper := 90.0
	bottom := -90.0
	return RequiredFloatValidators(key, FloatInRangeValidator(key, FloatRange{&upper, &bottom}))
}

func DistanceValidator(key string) Validator {
	var bottom int64 = 0
	var upper int64 = 32000000000
	return Int64InRangeValidator(key, Int64Range{Bottom: &bottom, Upper: &upper})
}
