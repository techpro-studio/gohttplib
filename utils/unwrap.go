package utils

func UnwrapOrDefaultInt(value *int, d int) int {
	if value != nil {
		return *value
	}
	return d
}

func UnwrapOrDefaultInt64(value *int64, d int64) int64 {
	if value != nil {
		return *value
	}
	return d
}

func UnwrapOrDefaultString(value *string, d string) string {
	if value != nil {
		return *value
	}
	return d
}

func UnwrapOrDefaultBool(value *bool, d bool) bool {
	if value != nil {
		return *value
	}
	return d
}

