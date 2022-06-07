package service

import (
	"gitlab.ozon.dev/hw/homework-2/internal/consts"
	"time"
)

// run background and after certain time collect expiredTimes
// then we notify users and delete this expiredTimes from map
func (s *GameServiceServer) garbageCollectorForTimes() {
	for {
		time.Sleep(time.Second * consts.SleepTime)
		if len(s.expTimes) == 0 {
			continue
		}
		now := time.Now()
		for expTime, resp := range s.expTimes {
			if expTime.Before(now) {
				// send notification that time is up
				if s.currPlayingChat[resp.ChatId] != "" {
					s.notificationChan <- resp
				}

				// delete from our map
				s.currPlayingChat[resp.ChatId] = ""
				delete(s.expTimes, expTime)
			}
		}
	}
}
