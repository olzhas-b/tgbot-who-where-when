package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/hw/homework-2/api"
	telegram2 "gitlab.ozon.dev/hw/homework-2/internal/client/telegram"
	"gitlab.ozon.dev/hw/homework-2/internal/client/validator"
	"gitlab.ozon.dev/hw/homework-2/internal/config"
	"golang.org/x/sync/semaphore"
	"google.golang.org/grpc"
	"log"
	"runtime"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		log.Fatalf("client.main.config.InitConfig got err %v", err)
	}
	log.Println("trying to connect")
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.HTTP.Name, cfg.HTTP.Port), opts...)
	if err != nil {
		log.Fatalf("client.main.grpc.Dial got err %v", err)
	}
	client := api.NewGameClient(conn)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)

	bot := telegram2.New().Bot
	updateCfg := telegram2.NewUpdateConfig()

	updates := bot.GetUpdatesChan(updateCfg)
	mxGoroutine := 2 * runtime.NumCPU()
	sem := semaphore.NewWeighted(int64(mxGoroutine))
	for update := range updates {
		if ok := validator.ValidateMessage(update.Message); !ok {
			continue
		}
		if err := sem.Acquire(context.Background(), 1); err != nil {
			continue
		}
		update := update
		go func() {
			defer sem.Release(1)
			go handleMessage(bot, &client, update)
		}()
	}
}

func handleMessage(bot *tgbotapi.BotAPI, client *api.GameClient, update tgbotapi.Update) {
	if bot == nil || client == nil {
		return
	}
	start, err := (*client).Start(context.Background(), &api.Message{Id: 1})
	if err != nil {
		return
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, start.Text)
	msg.ReplyToMessageID = update.Message.MessageID
	_, _ = bot.Send(msg)
}
