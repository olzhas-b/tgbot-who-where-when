package tools

import (
	"gitlab.ozon.dev/hw/homework-2/api"
	"gitlab.ozon.dev/hw/homework-2/internal/app/models"
	"strings"
)

func ConvertGrpcRespToDTO(resp *api.Response) models.Response {
	if resp == nil {
		return models.Response{}
	}
	return models.Response{
		Text:      resp.Text,
		MessageID: resp.MessageId,
		ChatID:    resp.ChatId,
		IsReply:   resp.IsReply,
	}
}

func ConvertGrpcReqToDTO(req *api.Request) models.Request {
	if req == nil {
		return models.Request{}
	}
	return models.Request{
		ChatID:   req.ChatId,
		UserName: req.UserName,
		UserID:   req.UserId,
		FullName: req.FullName,
	}
}

func ConvertToEscapedString(text string) string {
	var builder strings.Builder
	for _, r := range text {
		switch r {
		case '_', '*', '[', ']', '(', ')', '~', '`', '>', '#', '+', '-', '=', '|', '{', '}', '.', '!':
			builder.WriteString("\\")
			builder.WriteRune(r)
			break
		default:
			builder.WriteRune(r)
		}

	}
	return builder.String()
}
