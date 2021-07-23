package utils

// ReverseIntInt swap keys and values of map
func ReverseIntInt(dict map[int]int) map[int]int {
	reverseDict := make(map[int]int)

	for k, v := range dict {
		reverseDict[v] = k
	}
	return reverseDict
}

// ReverseIntString swap keys and values
func ReverseIntString(dict map[int]string) map[string]int {
	reverseDict := make(map[string]int)

	for k, v := range dict {
		reverseDict[v] = k
	}
	return reverseDict
}

// ReverseStringInt swap keys and values
func ReverseStringInt(dict map[string]int) map[int]string {
	reverseDict := make(map[int]string)

	for k, v := range dict {
		reverseDict[v] = k
	}
	return reverseDict
}

// ReverseStringString swap keys and values
func ReverseStringString(dict map[string]string) map[string]string {
	reverseDict := make(map[string]string)

	for k, v := range dict {
		reverseDict[v] = k
	}
	return reverseDict
}
