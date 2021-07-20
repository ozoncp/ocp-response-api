package domain

import "fmt"

type Response struct {
	Id        uint64
	UserId    uint64
	RequestId uint64
	Text      string
}

func NewResponse(id uint64, userId uint64, requestId uint64, text string) *Response {
	return &Response{
		Id:        id,
		UserId:    userId,
		RequestId: requestId,
		Text:      text,
	}
}

func (r Response) ToString() string {
	return fmt.Sprintf("Id: %d, UserId: %d, RequestId: %d, Text: %s", r.Id, r.UserId, r.RequestId, r.Text)
}
