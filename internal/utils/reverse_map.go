package utils

import "fmt"

// ReverseIntInt swap keys and values of map
func ReverseIntInt(dict map[int]int) (map[int]int, error) {
	reverseDict := make(map[int]int)

	for k, v := range dict {
		if _, ok := reverseDict[v]; ok {
			return nil, fmt.Errorf("error while reversing map: dublicate value was found, value: [%d]:[%d]", k, v)
		}
		reverseDict[v] = k
	}
	return reverseDict, nil
}

// ReverseIntString swap keys and values
func ReverseIntString(dict map[int]string) (map[string]int, error) {
	reverseDict := make(map[string]int)

	for k, v := range dict {
		if _, ok := reverseDict[v]; ok {
			return nil, fmt.Errorf("error while reversing map: dublicate value was found, value: [%d]:[%s]", k, v)
		}
		reverseDict[v] = k
	}
	return reverseDict, nil
}

// ReverseStringInt swap keys and values
func ReverseStringInt(dict map[string]int) (map[int]string, error) {
	reverseDict := make(map[int]string)

	for k, v := range dict {
		if _, ok := reverseDict[v]; ok {
			return nil, fmt.Errorf("error while reversing map: dublicate value was found, value: [%s]:[%d]", k, v)
		}
		reverseDict[v] = k
	}
	return reverseDict, nil
}

// ReverseStringString swap keys and values
func ReverseStringString(dict map[string]string) (map[string]string, error) {
	reverseDict := make(map[string]string)

	for k, v := range dict {
		if _, ok := reverseDict[v]; ok {
			return nil, fmt.Errorf("error while reversing map: dublicate value was found, value: [%s]:[%s]", k, v)
		}
		reverseDict[v] = k
	}
	return reverseDict, nil
}
