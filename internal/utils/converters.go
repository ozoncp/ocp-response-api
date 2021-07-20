package utils

import (
	"errors"
	"github.com/ozoncp/ocp-response-api/domain"
)

// ToMap Create map from slice of Responses
func ToMap(entities []domain.Response) (map[uint64]domain.Response, error) {
	dict := make(map[uint64]domain.Response)
	for _, v := range entities {
		if _, ok := dict[v.Id]; ok {
			return nil, errors.New(DuplicateIdError)
		}
		dict[v.Id] = v
	}

	return dict, nil
}
