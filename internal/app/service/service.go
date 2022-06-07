package service

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"gitlab.ozon.dev/hw/homework-2/api"
	"gitlab.ozon.dev/hw/homework-2/internal/app/models"
	"gitlab.ozon.dev/hw/homework-2/internal/app/repository"
	"gitlab.ozon.dev/hw/homework-2/internal/consts"
	"gitlab.ozon.dev/hw/homework-2/tools"
	"io"
	"log"
	"strings"
	"sync"
	"time"
)

type GameServiceServer struct {
	mu               sync.RWMutex
	repo             *repository.Repository
	expTimes         map[time.Time]*api.Response
	currPlayingChat  map[int64]models.Answer
	notificationChan chan *api.Response
	api.UnimplementedGameServer
}

func NewGameServiceServer(repo *repository.Repository) *GameServiceServer {
	server := GameServiceServer{
		currPlayingChat:  make(map[int64]models.Answer),
		expTimes:         make(map[time.Time]*api.Response),
		notificationChan: make(chan *api.Response, 10000),
		repo:             repo,
	}
	go server.garbageCollectorForTimes()
	return &server
}

func (srv *GameServiceServer) Start(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	if ok, _ := srv.currPlayingChat[req.ChatId]; len(ok) != 0 {
		return &api.Response{Text: "Вы уже начали игру", ChatId: req.ChatId}, nil
	}
	var question, answer string
	defer func() {
		if err == nil {
			srv.currPlayingChat[req.ChatId] = models.Answer(answer)
			srv.expTimes[time.Now().Add(time.Second*consts.WaitingAnswer)] =
				&api.Response{
					Text:   fmt.Sprintf("%s \nВопрос: %s \nОтвет: ||%s||", consts.TimeIsUpMsg, tools.ConvertToEscapedString(question), tools.ConvertToEscapedString(answer)),
					ChatId: req.ChatId,
				}
		}
	}()
	question, answer, err = srv.repo.IGameRepostiory.GetRandomQuestion(ctx)
	if err != nil {
		return
	}
	return &api.Response{Text: tools.ConvertToEscapedString(question), ChatId: req.ChatId}, nil
}

func (srv *GameServiceServer) CheckAnswer(ctx context.Context, req *api.Request) (*api.Response, error) {
	if len(req.UserAnswer) == 0 {
		return &api.Response{
			Text:   consts.DontStartGameMsg,
			ChatId: req.ChatId,
		}, nil
	}
	if strings.ToLower(req.UserAnswer) == strings.ToLower(string(srv.currPlayingChat[req.ChatId])) {
		srv.mu.Lock()
		delete(srv.currPlayingChat, req.ChatId)
		srv.mu.Unlock()
		if err := srv.repo.IGameRepostiory.AddPointToUser(ctx, tools.ConvertGrpcReqToDTO(req)); err != nil {
			log.Println("GameServiceServer.CheckAnswer.AddPointToUser got error", err)
		}
		return &api.Response{
			Text:      tools.ConvertToEscapedString(fmt.Sprintf(consts.CongratulationMsg, req.FullName)),
			ChatId:    req.ChatId,
			MessageId: req.MessageId,
		}, nil
	}
	return &api.Response{Text: consts.WrongAnswerMsg, ChatId: req.ChatId}, nil
}

func (srv *GameServiceServer) GetTop10Players(ctx context.Context, req *api.Request) (*api.Response, error) {
	var (
		users models.Users
		err   error
	)
	users, err = srv.repo.IGameRepostiory.GetLeaderboards(ctx, req.ChatId)
	if err != nil {
		return nil, err
	}
	return &api.Response{
		Text:   tools.ConvertToEscapedString(users.ConvertToString()),
		ChatId: req.ChatId,
	}, nil
}

func (srv *GameServiceServer) Notification(emp *empty.Empty, notification api.Game_NotificationServer) error {
	for resp := range srv.notificationChan {
		if err := notification.Send(resp); err != nil {
			if err == io.EOF {
				return nil
			}
			log.Println(err)
		}
	}
	return nil
}
