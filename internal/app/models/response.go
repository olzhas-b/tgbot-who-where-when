package models

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Response struct {
	Text      string
	ChatID    int64
	MessageID int64
	IsReply   bool
}

func (resp *Response) ConvertToMsg() tgbotapi.MessageConfig {
	return tgbotapi.MessageConfig{
		Text: resp.Text,
		BaseChat: tgbotapi.BaseChat{
			ChatID:           resp.ChatID,
			ReplyToMessageID: int(resp.MessageID),
		},
		DisableWebPagePreview: false,
		ParseMode:             tgbotapi.ModeMarkdownV2,
	}
}
