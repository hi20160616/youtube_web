package json

import (
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/api/youtube/v3"
)

type VideosParams struct {
	Part string
	Id   string
}

func (vp *VideosParams) List() (*youtube.VideoListResponse, error) {
	if vp.Part == "" {
		vp.Part = "snippet, contentDetails"
	}
	if vp.Id == "" {
		return nil, errors.New("db: VideosParams List: id is nil")
	}

	call := YoutubeService.Videos.List(strings.Split(vp.Part, ","))
	call = call.Id(strings.Split(vp.Id, ",")...)
	return call.Do()
}
