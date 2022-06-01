package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/hw/homework-2/internal/config"
	"log"
)

const (
	token = "5568592349:AAFoLMsEza6rzyAuTQDsTW9Um2ZRCMX4H90"
)

type Telegram struct {
	Bot *tgbotapi.BotAPI
}

func New() Telegram {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		return Telegram{}
	}

	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Fatalln("create NewBotAPI got error %v", err)
	}
	if cfg.Telegram.Debug {
		bot.Debug = true
	}
	return Telegram{Bot: bot}
	//u := tgbotapi.NewUpdate(0)

	//updates := bot.GetUpdatesChan(u)
	//for update := range updates {
	//	if update.Message != nil {
	//		fmt.Printf("user %s send message: %s\n", update.Message.From.UserName, update.Message.Text)
	//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "wtf")
	//		msg.ReplyToMessageID = update.Message.MessageID
	//		_, _ = bot.Send(msg)
	//
	//	}
	//}
}
