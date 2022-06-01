package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewUpdateConfig() tgbotapi.UpdateConfig {
	updateCfg := tgbotapi.NewUpdate(0)
	return updateCfg
}
