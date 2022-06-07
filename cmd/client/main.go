package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/protobuf/ptypes/empty"
	"gitlab.ozon.dev/hw/homework-2/api"
	"gitlab.ozon.dev/hw/homework-2/internal/app/models"
	tg "gitlab.ozon.dev/hw/homework-2/internal/client/telegram"
	"gitlab.ozon.dev/hw/homework-2/internal/client/validator"
	"gitlab.ozon.dev/hw/homework-2/internal/config"
	"gitlab.ozon.dev/hw/homework-2/internal/consts"
	"gitlab.ozon.dev/hw/homework-2/tools"
	"golang.org/x/sync/semaphore"
	"google.golang.org/grpc"
	"io"
	"log"
	"runtime"
	"time"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	// read config from config.yaml
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		log.Fatalf("client.main.config.InitConfig got err %v", err)
	}
	log.Println("trying to connect grpcServer")
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.HTTP.Name, cfg.HTTP.Port), opts...)
	if err != nil {
		log.Fatalf("client.main.grpc.Dial got err %v", err)
	}
	client := api.NewGameClient(conn)

	// create bot
	bot := tg.NewBot()
	updateCfg := tg.NewUpdateConfig()
	updates := bot.GetUpdatesChan(updateCfg)

	// handle notification from server
	notif, err := client.Notification(context.Background(), &empty.Empty{})
	if err != nil {
		log.Println(err)
	}
	go handleNotification(bot, notif)

	// limited max goroutine with semaphore
	mxGoroutine := 4 * runtime.NumCPU()
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
			ctx, _ := context.WithTimeout(context.Background(), time.Minute*30)
			defer sem.Release(1)
			handleMessage(ctx, bot, &client, update)
		}()
	}
}

func handleMessage(ctx context.Context, bot *tg.Bot, client *api.GameClient, update tgbotapi.Update) {
	request, err := parseUpdate(update)
	if err != nil {
		return
	}
	switch request.Command {
	case consts.Start:
		resp, err := (*client).Start(ctx, request.ConvertToGrpcMsg())
		if err != nil || resp == nil {
			log.Println("answer got error: ", err)
			return
		}
		bot.Send(tools.ConvertGrpcRespToDTO(resp))
	case consts.Answer:
		if len(request.UserAnswer) == 0 {
			break
		}
		resp, err := (*client).CheckAnswer(ctx, request.ConvertToGrpcMsg())
		if err != nil || resp == nil {
			log.Println("answer got error: ", err)
			return
		}
		bot.Send(tools.ConvertGrpcRespToDTO(resp))
	case consts.Top:
		resp, err := (*client).GetTop10Players(ctx, request.ConvertToGrpcMsg())
		if err != nil || resp == nil {
			log.Println("answer got error: ", err)
			return
		}
		bot.Send(tools.ConvertGrpcRespToDTO(resp))
	case consts.Help:
		bot.Send(models.Response{
			ChatID: request.ChatID,
			Text:   tools.ConvertToEscapedString(consts.Help),
		})
	}
}

func parseUpdate(update tgbotapi.Update) (models.Request, error) {
	request := models.Request{
		UserName:   update.Message.From.UserName,
		UserID:     update.Message.From.ID,
		FullName:   fmt.Sprintf("%s %s", update.Message.From.FirstName, update.Message.From.LastName),
		UserAnswer: update.Message.CommandArguments(),
		ChatID:     update.Message.Chat.ID,
		Command:    update.Message.Command(),
		MessageID:  int64(update.Message.MessageID),
		IsCommand:  update.Message.IsCommand(),
	}

	return request, nil
}

func handleNotification(bot *tg.Bot, notif api.Game_NotificationClient) {
	go func() {
		for {
			msg, err := notif.Recv()
			if err == io.EOF {
				return
			}
			if err != nil || msg == nil {
				continue
			}
			bot.Send(tools.ConvertGrpcRespToDTO(msg))
		}
	}()
}
