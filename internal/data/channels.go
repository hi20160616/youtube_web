package data

import (
	db "github.com/hi20160616/youtube_web/internal/pkg/db/json"
	"google.golang.org/api/youtube/v3"
)

func ReadChannels() ([]*db.Channel, error) {
	return db.ReadChannels()
}

func ReadActivities() ([]*youtube.VideoListResponse, error) {
	return db.ReadActivities()
}
