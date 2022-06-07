package main

import (
	"context"
	"gitlab.ozon.dev/hw/homework-2/internal/app/repository"
	"gitlab.ozon.dev/hw/homework-2/internal/config"
	db2 "gitlab.ozon.dev/hw/homework-2/internal/db"
	"golang.org/x/sync/semaphore"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
)

func main() {
	cfg, _ := config.InitConfig("config.yaml")
	db := db2.InitDB(cfg)
	repo := repository.NewGameRepository(db)

	file, err := os.Open("questions.txt")
	if err != nil {
		log.Fatalf("questions.txt open got err :%v", err)
	}
	raw, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("readAll got error: %v", err)
	}
	lines := strings.Split(string(raw), "\n")
	mxGoroutine := runtime.NumCPU()
	sem := semaphore.NewWeighted(int64(2 * mxGoroutine))
	for _, line := range lines {
		if err := sem.Acquire(context.Background(), 1); err != nil {
			log.Fatalf("Acuire problem")
		}
		line := line
		go func() {
			defer sem.Release(1)
			words := strings.Split(line, "|")
			if len(words) != 2 {
				return
			}
			answer, question := words[0], words[1]
			if err := repo.InsertQuestion(context.Background(), answer, question); err != nil {
				log.Println(err)
			}
		}()
	}
}
