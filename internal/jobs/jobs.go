package jobs

import (
	"google.golang.org/api/youtube/v3"

	db "github.com/hi20160616/youtube_web/internal/pkg/db/json"
)

func UpdateChannels(s *youtube.Service) ([]*db.Channel, error) {
	return db.UpdateChannels()
}
