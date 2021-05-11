package jobs

import (
	"context"
	"log"
	"time"

	"google.golang.org/api/youtube/v3"

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
		log.Println("Update Channels ...")
		if _, err := UpdateChannels(); err != nil {
			return err
		}
		log.Println("Done.")
		log.Println("Update Activities ...")
		if _, err := UpdateActivities(); err != nil {
			return err
		}
		log.Println("Done.")
		return nil
	}

	for {
		select {
		case <-done:
			return errors.New("Exit Update on purpose!")
		case <-ctx.Done():
			return ctx.Err()
		default:
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
