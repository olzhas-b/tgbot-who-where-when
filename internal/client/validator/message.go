package validator

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/hw/homework-2/internal/consts"
	"strings"
)

func ValidateMessage(message *tgbotapi.Message) bool {
	if message == nil {
		return false
	}
	if !message.IsCommand() {
		return false
	}
	if !strings.Contains(message.CommandWithAt(), consts.BotName) {
		return false
	}
	return true
}
