package utils

func ReverseIntInt(dict map[int]int) map[int]int {
	reverseDict := make(map[int]int)

	for k, v := range dict {
		reverseDict[v] = k
	}
	return reverseDict
}

func ReverseIntString(dict map[int]string) map[string]int {
	reverseDict := make(map[string]int)

	for k, v := range dict {
		reverseDict[v] = k
	}
	return reverseDict
}

func ReverseStringInt(dict map[string]int) map[int]string {
	reverseDict := make(map[int]string)

	for k, v := range dict {
		reverseDict[v] = k
	}
	return reverseDict
}

func ReverseStringString(dict map[string]string) map[string]string {
	reverseDict := make(map[string]string)

	for k, v := range dict {
		reverseDict[v] = k
	}
	return reverseDict
}

func ReverseStringInterface(dict map[string]interface{}) map[interface{}]string {
	reverseDict := make(map[interface{}]string)

	for k, v := range dict {
		reverseDict[v] = k
	}
	return reverseDict
}

func ReverseInterfaceString(dict map[interface{}]string) map[string]interface{} {
	reverseDict := make(map[string]interface{})

	for k, v := range dict {
		reverseDict[v] = k
	}
	return reverseDict
}

func Reverse(dict map[interface{}]interface{}) map[interface{}]interface{} {
	reverseDict := make(map[interface{}]interface{})

	for k, v := range dict {
		reverseDict[v] = k
	}
	return reverseDict
}

func ReverseIntInterface(dict map[int]interface{}) map[interface{}]int {
	reverseDict := make(map[interface{}]int)

	for k, v := range dict {
		reverseDict[v] = k
	}
	return reverseDict
}

func ReverseInterfaceInt(dict map[interface{}]int) map[int]interface{} {
	reverseDict := make(map[int]interface{})

	for k, v := range dict {
		reverseDict[v] = k
	}
	return reverseDict
}
