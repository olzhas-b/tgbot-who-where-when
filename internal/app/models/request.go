package models

import "gitlab.ozon.dev/hw/homework-2/api"

type Request struct {
	UserName   string
	FullName   string
	Command    string
	UserAnswer string
	ChatID     int64
	UserID     int64
	MessageID  int64
	IsCommand  bool
}

func (req *Request) ConvertToGrpcMsg() *api.Request {
	return &api.Request{
		UserName:   req.UserName,
		FullName:   req.FullName,
		UserAnswer: req.UserAnswer,
		ChatId:     req.ChatID,
		MessageId:  req.MessageID,
		UserId:     req.UserID,
	}
}
