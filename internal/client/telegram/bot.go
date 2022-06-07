package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/hw/homework-2/internal/config"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot() *Bot {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		return &Bot{}
	}

	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Fatalf("create NewBotAPI got error %v", err)
	}
	if cfg.Telegram.Debug {
		bot.Debug = true
	}
	return &Bot{bot: bot}
}

func (tg *Bot) GetUpdatesChan(cfg tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel {
	return tg.bot.GetUpdatesChan(cfg)
}
