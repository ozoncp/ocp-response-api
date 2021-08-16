package saver_test

import (
	"github.com/ozoncp/ocp-response-api/internal/models"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"

	"github.com/ozoncp/ocp-response-api/internal/mocks"
	"github.com/ozoncp/ocp-response-api/internal/saver"

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
		{Id: 7, UserId: 7, RequestId: 7, Text: "text7"},
		{Id: 8, UserId: 8, RequestId: 8, Text: "text8"},
		{Id: 9, UserId: 9, RequestId: 9, Text: "text9"},
		{Id: 10, UserId: 10, RequestId: 10, Text: "text10"},
	}
}

var _ = Describe("Saver", func() {
	var (
		ctrl    *gomock.Controller
		requests    []models.Response
		flusher *mocks.MockFlusher
		sav         saver.Saver
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		flusher = mocks.NewMockFlusher(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("", func() {
		JustBeforeEach(func() {
			sav = saver.NewSaver(10, flusher, 3 * time.Second)
			requests = getResponses()
		})

		It("", func() {
			defer sav.Close()

			flusher.EXPECT().
				Flush(requests).
				Return(nil, nil).
				MaxTimes(1).
				MinTimes(1)

			for _, req := range requests {
				sav.Save(req)
			}
		})

		It("", func() {
			defer sav.Close()
			flusher.EXPECT().
				Flush(requests).
				Return(nil, nil).
				MaxTimes(1).
				MinTimes(1)

			for _, req := range requests[:5] {
				sav.Save(req)
			}
			time.Sleep(time.Second)
			for _, req := range requests[5:] {
				sav.Save(req)
			}
		})
	})
})
