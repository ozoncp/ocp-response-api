package utils

// FilterInt filters slice of ints
func FilterInt(slice []int, filterFunc func(int) bool) []int {
	var filtered []int
	for _, item := range slice {
		if filterFunc(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

// FilterString filters slice of strings
func FilterString(slice []string, filterFunc func(string) bool) []string {
	var filtered []string
	for _, item := range slice {
		if filterFunc(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
