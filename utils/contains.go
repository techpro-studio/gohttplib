package utils

func ContainsString(array []string, element string) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}

func ContainsInt(array []int, element int) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}
