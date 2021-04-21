package data

import (
	db "github.com/hi20160616/youtube_web/internal/pkg/db/json"
	"google.golang.org/api/youtube/v3"
)

type VideosRepo struct {
	vs *db.Videos
}

func NewVideoRepo() *VideosRepo {
	return &VideosRepo{&db.Videos{}}
}

func (vr *VideosRepo) SetChannelId(cid string) *VideosRepo {
	vr.vs.ChannelId = cid
	return vr
}

func (vr *VideosRepo) SetMaxResults(max int64) *VideosRepo {
	vr.vs.MaxResults = max
	return vr
}

func (vr *VideosRepo) GetVideos() (*youtube.ActivityListResponse, error) {
	return vr.vs.GetVideos()
}
