package utils

func FilterInt(slice []int, f func(int) bool) []int {
	var filtered []int
	for _, item := range slice {
		if f(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func FilterString(slice []string, f func(string) bool) []string {
	var filtered []string
	for _, item := range slice {
		if f(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func FilterFloat32(slice []float32, f func(float32) bool) []float32 {
	var filtered []float32
	for _, item := range slice {
		if f(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func FilterFloat64(slice []float64, f func(float64) bool) []float64 {
	var filtered []float64
	for _, item := range slice {
		if f(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
