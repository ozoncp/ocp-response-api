package flusher_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	"github.com/onsi/gomega"

	"github.com/ozoncp/ocp-response-api/internal/flusher"
	"github.com/ozoncp/ocp-response-api/internal/mocks"
	"github.com/ozoncp/ocp-response-api/internal/models"
)

func getResponses() []models.Response {
	return []models.Response{
		{Id: 0, UserId: 0, RequestId: 0, Text: "text0"},
		{Id: 1, UserId: 1, RequestId: 1, Text: "text1"},
		{Id: 2, UserId: 2, RequestId: 2, Text: "text2"},
		{Id: 3, UserId: 3, RequestId: 3, Text: "text3"},
		{Id: 4, UserId: 4, RequestId: 4, Text: "text4"},
		{Id: 5, UserId: 5, RequestId: 5, Text: "text5"},
		{Id: 6, UserId: 6, RequestId: 6, Text: "text6"},
	}
}

var _ = Describe("Flusher", func() {
	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo
		f        flusher.Flusher
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		f = flusher.NewFlusher(5, mockRepo)
	})

	AfterEach(func() { ctrl.Finish() })

	table.DescribeTable("Test flusher",
		func(responses []models.Response, repoReturns error, expected interface{}) {
			mockRepo.EXPECT().
				AddResponses(gomock.Any()).
				Return(repoReturns).MaxTimes(2)

			gomega.Ω(f.Flush(responses)).To(gomega.Equal(expected))
		},

		table.Entry("ok", getResponses(), nil, getResponses()),
		table.Entry("error", getResponses(), errors.New("test"), getResponses()))
})
