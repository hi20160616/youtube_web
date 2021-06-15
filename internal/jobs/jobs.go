package jobs

import (
	"context"
	"log"
	"time"

	"google.golang.org/api/youtube/v3"

	"github.com/hi20160616/youtube_web/configs"
	db "github.com/hi20160616/youtube_web/internal/pkg/db/json"
	"github.com/pkg/errors"
)

var (
	done = make(chan struct{}, 1)
	sema = make(chan struct{}, 1)
)

func UpdateChannels() ([]*db.Channel, error) {
	// log.Println("UpdateChannels biz test") // for test
	// return nil, nil                        // for test
	return db.UpdateChannels()
}

func UpdateActivities() ([]*youtube.VideoListResponse, error) {
	// log.Println("UpdateActivities biz test") // for test
	// return nil, nil                          // for test
	return db.UpdateActivities()
}

func UpdateByHoursStart(ctx context.Context) error {
	doit := func() error {
		log.Printf("[%s] Update Channels ...", configs.Value.Title)
		if _, err := UpdateChannels(); err != nil {
			if errors.Is(err, context.Canceled) {
				return err
			}
			log.Printf("[%s] UpdateByHoursStart: %v", configs.Value.Title, err)
		}
		log.Printf("[%s] Update Channels Done.", configs.Value.Title)
		log.Printf("[%s] Update Activities ...", configs.Value.Title)
		if _, err := UpdateActivities(); err != nil {
			if errors.Is(err, context.Canceled) {
				return err
			}
			log.Printf("[%s] UpdateByHoursStart: %v", configs.Value.Title, err)
		}
		log.Printf("[%s] Update Activities Done.", configs.Value.Title)
		return nil
	}

	// Run once while just start
	doit()

	for {
		select {
		case <-done:
			return errors.Errorf("[%s] Exit Update on purpose!", configs.Value.Title)
		case <-ctx.Done():
			return ctx.Err()
		case <-time.Tick(time.Second):
			if time.Now().Minute() == 0 && time.Now().Second() == 0 {
				sema <- struct{}{}
				doit()
				<-sema
			}
		}
	}
}

func UpdateByHoursStop() error {
	close(done)
	return nil
}
