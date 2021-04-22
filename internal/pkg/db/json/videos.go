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

func (vp *VideosParams) List(id string) (*youtube.VideoListResponse, error) {
	if vp.Part == "" {
		vp.Part = "snippet, contentDetails"
	}
	if vp.Id = id; id == "" {
		return nil, errors.New("db: json: id is nil")
	}

	call := YoutubeService.Videos.List(strings.Split(vp.Part, ","))
	call = call.Id(strings.Split(id, ",")...)
	return call.Do()
}
