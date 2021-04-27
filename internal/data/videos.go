package data

import (
	"strings"

	db "github.com/hi20160616/youtube_web/internal/pkg/db/json"
	"google.golang.org/api/youtube/v3"
)

type VideosRepo struct {
	Activities *db.ActivitiesParams
	Videos     *db.VideosParams
}

func NewVideosRepo() *VideosRepo {
	return &VideosRepo{&db.ActivitiesParams{}, &db.VideosParams{}}
}

func (vr *VideosRepo) WithChannelId(cid string) *VideosRepo {
	vr.Activities.ChannelId = cid
	return vr
}

func (vr *VideosRepo) WithMaxResults(max int64) *VideosRepo {
	vr.Activities.MaxResults = max
	return vr
}

func (vr *VideosRepo) WithPublishAfter4Activities(minutes int) *VideosRepo {
	vr.Activities = vr.Activities.WithPublishAfter(minutes)
	return vr
}

func (vr *VideosRepo) GetActivitiesVideos() (*youtube.VideoListResponse, error) {
	vids := []string{}
	a_vs, err := vr.Activities.List()
	if err != nil {
		return nil, err
	}
	for _, v := range a_vs.Items {
		if v.ContentDetails.Upload != nil {
			vids = append(vids, v.ContentDetails.Upload.VideoId)
		}
	}
	return vr.Videos.List(strings.Join(vids, ","))
}

// GetVideos return all videos by id
func (vr *VideosRepo) GetVideos() (*youtube.VideoListResponse, error) {
	return vr.Videos.List(vr.Videos.Id)
}

// WithVideoId only used for GetVideos()
func (vr *VideosRepo) WithVideoId(vid string) *VideosRepo {
	vr.Videos.Id = vid
	return vr
}
