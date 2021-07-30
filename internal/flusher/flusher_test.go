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
	}
}

var _ = Describe("Flusher", func() {
	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockResponseRepo
		f        flusher.Flusher
		pMarker  string
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockResponseRepo(ctrl)
		f = flusher.NewFlusher(5, mockRepo)
		pMarker = "panic"
	})

	AfterEach(func() { ctrl.Finish() })

	table.DescribeTable("Test flusher",
		func(responses []models.Response, repoReturns error, expected interface{}) {
			mockRepo.EXPECT().
				AddResponses(gomock.Any()).
				Return(repoReturns)

			if expected == pMarker {
				gomega.Expect(func() { f.Flush(responses) }).Should(gomega.Panic())
			} else {
				gomega.Ω(f.Flush(responses)).To(gomega.Equal(expected))
			}
		},

		table.Entry("ok", getResponses(), nil, getResponses()),
		table.Entry("panic", getResponses(), errors.New("test"), pMarker))

})
