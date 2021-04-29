package jobs

import (
	"context"
	"log"
	"time"

	"google.golang.org/api/youtube/v3"

	db "github.com/hi20160616/youtube_web/internal/pkg/db/json"
	"github.com/pkg/errors"
)

var Done = make(chan struct{}, 1)

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
		// if time.Now().Second() == 0 { // for test
		//         time.Sleep(1 * time.Second) // for test
		if time.Now().Minute() == 0 {
			log.Println("Update Channels ...")
			if _, err := UpdateChannels(); err != nil {
				return err
			}
			log.Println("Update Activities ...")
			if _, err := UpdateActivities(); err != nil {
				return err
			}
		}
		return nil
	}
	for {
		select {
		case <-Done:
			return errors.New("Exit Update on purpose!")
		case <-ctx.Done():
			return ctx.Err()
		default:
			doit()
		}
	}
}

func UpdateByHoursStop() error {
	close(Done)
	return nil
}
