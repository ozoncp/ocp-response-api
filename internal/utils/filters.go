package utils

import "github.com/ozoncp/ocp-response-api/internal/models"

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

// FilterResponse filters slice of Response
func FilterResponse(slice []models.Response, filterFunc func(models.Response) bool) []models.Response {
	var filtered []models.Response
	for _, item := range slice {
		if filterFunc(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
