package tests

import (
	"github.com/ozoncp/ocp-response-api/domain"
	"github.com/ozoncp/ocp-response-api/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToMap(t *testing.T) {
	responses := Get10Responses()

	dict, err := utils.ToMap(responses)

	assert.Equal(t, err, nil)
	assert.NotEqual(t, dict, nil)
	assert.Equal(t, len(dict), len(responses))

	for _, resp := range responses {
		assert.Equal(t, dict[resp.Id], resp)
	}
}

func TestToMapWithError(t *testing.T) {
	responses := Get10Responses()
	responses = append(responses, *domain.NewResponse(1, 1, 1, "test1"))
	_, err := utils.ToMap(responses)

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), utils.DuplicateIdError)
}
