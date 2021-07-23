package utils

import (
	"fmt"
)

// ToMap create map from slice of Responses
func ToMap(entities []Response) (map[uint64]Response, error) {
	res := make(map[uint64]Response, len(entities))
	for _, entity := range entities {
		if _, ok := res[entity.Id]; ok {
			return nil, fmt.Errorf("error while transofming responses to map: dublicate response was found: [%s]", entity.ToString())
		}
		res[entity.Id] = entity
	}

	return res, nil
}
