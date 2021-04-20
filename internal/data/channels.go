package data

import (
	db "github.com/hi20160616/youtube_web/internal/pkg/db/json"
)

func ReadChannels() ([]*db.Channel, error) {
	return db.ReadChannels()
}
