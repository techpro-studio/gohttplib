package utils


func AsArrayOfString(arr []interface{}) []string {
	typed := []string{}
	for _, item := range arr {
		typed = append(typed, item.(string))
	}
	return typed
}


func AsArrayOfInt(arr []interface{}) []int {
	typed := []int{}
	for _, item := range arr {
		typed = append(typed, item.(int))
	}
	return typed
}
