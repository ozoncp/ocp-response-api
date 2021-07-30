package models

import "fmt"

type Response struct {
	Id        uint64
	UserId    uint64
	RequestId uint64
	Text      string
}

func (r Response) ToString() string {
	return fmt.Sprintf("%#v", r)
}
