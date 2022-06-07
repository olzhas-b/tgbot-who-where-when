package telegram

import (
	"gitlab.ozon.dev/hw/homework-2/internal/app/models"
	"log"
)

func (b *Bot) Send(resp models.Response) {
	msg := resp.ConvertToMsg()
	res, err := b.bot.Send(msg)
	if err != nil {
		log.Printf("bot.Send body[%v], got error: %v", res, err)
	}
}
