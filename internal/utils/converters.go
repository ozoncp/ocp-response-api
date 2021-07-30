package utils

import (
	"fmt"
	"github.com/ozoncp/ocp-response-api/internal/models"
)

// ToMap create map from slice of Responses
func ToMap(entities []models.Response) (map[uint64]models.Response, error) {
	res := make(map[uint64]models.Response, len(entities))
	for _, entity := range entities {
		if _, ok := res[entity.Id]; ok {
			return nil, fmt.Errorf("error while transofming responses to map: dublicate response.id was found: [%s]", entity.ToString())
		}
		res[entity.Id] = entity
	}

	return res, nil
}
