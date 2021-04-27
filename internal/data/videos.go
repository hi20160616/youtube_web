package data

import (
	"strings"

	db "github.com/hi20160616/youtube_web/internal/pkg/db/json"
	"github.com/pkg/errors"
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

func (vr *VideosRepo) List() (*youtube.VideoListResponse, error) {
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
	vr.Videos.Id = strings.Join(vids, ",")
	return vr.Videos.List()
}

// GetVideos return all videos by id
func (vr *VideosRepo) GetVideos() (*youtube.VideoListResponse, error) {
	if vr.Videos.Id == "" {
		return nil, errors.New("data: GetVideos by nil id")
	}
	return vr.Videos.List()
}

// WithVideoId only used for GetVideos()
func (vr *VideosRepo) WithVideoId(vid string) *VideosRepo {
	vr.Videos.Id = vid
	return vr
}
