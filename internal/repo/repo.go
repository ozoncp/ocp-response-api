package repo

import "github.com/ozoncp/ocp-response-api/internal/models"

type ResponseRepo interface {
	AddResponses(responses []models.Response) error
	RemoveResponse(responseId uint64) error
	DescribeResponse(responseId uint64) (*models.Response, error)
	ListResponses(limit, offset uint64) ([]models.Response, error)
}
