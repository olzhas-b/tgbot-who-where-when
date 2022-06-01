package validator

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func ValidateMessage(message *tgbotapi.Message) bool {
	if message == nil {
		return false
	}
	return true
}
