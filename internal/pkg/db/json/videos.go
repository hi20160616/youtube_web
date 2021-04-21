package json

import (
	"strings"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/api/youtube/v3"
)

type Videos struct {
	Part          string // should be snippet,contentDetails
	ChannelId     string
	ChannelTitle  string
	MaxResults    int64
	PageToken     string
	NextPageToken string
	PublishAfter  string
	PublishBefore string
}

func (v *Videos) GetVideos() (*youtube.ActivityListResponse, error) {
	timeParsed := func(minutes int) string {
		return time.Now().Add(time.Duration(minutes) * time.Minute).Format(time.RFC3339)
	}

	// init parameters
	if v.Part == "" {
		v.Part = "snippet,contentDetails"
	}
	if v.ChannelId == "" {
		return nil, errors.New("data: GetVideos: ChannelId is nil")
	}
	if v.MaxResults == 0 {
		v.MaxResults = 10
	}
	if v.MaxResults >= 50 {
		v.MaxResults = 50
	}
	if v.PublishAfter == "" {
		v.PublishAfter = timeParsed(-1 * 60 * 24) // 1 day ago
	}
	if v.PublishBefore == "" {
		v.PublishBefore = timeParsed(0)
	}

	// load
	call := YoutubeService.Activities.List(strings.Split(v.Part, ","))
	call = call.ChannelId(v.ChannelId)
	call = call.MaxResults(v.MaxResults)
	call = call.PublishedAfter(v.PublishAfter)
	call = call.PublishedBefore(v.PublishBefore)
	if v.NextPageToken != "" {
		call = call.PageToken(v.NextPageToken)
	}

	// do it
	res, err := call.Do()
	if err != nil {
		return nil, errors.WithMessage(err, "data: GetVideos error: ")
	}
	if len(res.Items) > 0 {
		v.ChannelTitle = res.Items[0].Snippet.Title
	}
	v.NextPageToken = res.NextPageToken
	return res, err
}
