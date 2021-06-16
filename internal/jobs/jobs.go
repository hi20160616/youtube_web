package jobs

import (
	"context"
	"log"

	"google.golang.org/api/youtube/v3"

	db "github.com/hi20160616/youtube_web/internal/pkg/db/json"
	"github.com/pkg/errors"
)

func UpdateChannels() ([]*db.Channel, error) {
	return db.UpdateChannels()
}

func UpdateActivities() ([]*youtube.VideoListResponse, error) {
	return db.UpdateActivities()
}

func UpdateActions(ctx context.Context) error {
	log.Printf("Update Channels ...")
	if _, err := UpdateChannels(); err != nil {
		if errors.Is(err, context.Canceled) {
			return err
		}
		log.Printf("UpdateChannels uncomplete: %v", err)
	}
	log.Printf("Update Channels Done.")
	log.Printf("Update Activities ...")
	if _, err := UpdateActivities(); err != nil {
		if errors.Is(err, context.Canceled) {
			return err
		}
		log.Printf("UpdateActivities uncomplete: %v", err)
	}
	log.Printf("Update Activities Done.")
	return ctx.Err()
}
