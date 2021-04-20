package data

import (
	"strings"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/api/youtube/v3"
)

type Videos struct {
	part          string // should be snippet,contentDetails
	channelId     string
	channelTitle  string
	maxResults    int64
	pageToken     string
	nextPageToken string
	publishAfter  string
	publishBefore string
}

func (v *Videos) GetVideos(s *youtube.Service) (*youtube.ActivityListResponse, error) {
	timeParsed := func(minutes int) string {
		return time.Now().Add(time.Duration(minutes) * time.Minute).Format(time.RFC3339)
	}

	// init parameters
	if v.part == "" {
		v.part = "snippet,contentDetails"
	}
	if v.channelId == "" {
		return nil, errors.New("data: GetVideos: channelId is nil")
	}
	if v.maxResults == 0 {
		v.maxResults = 10
	}
	if v.maxResults >= 50 {
		v.maxResults = 50
	}
	if v.publishAfter == "" {
		v.publishAfter = timeParsed(-1 * 60 * 24) // 1 day ago
	}
	if v.publishBefore == "" {
		v.publishBefore = timeParsed(0)
	}

	// load
	call := s.Activities.List(strings.Split(v.part, ","))
	call = call.ChannelId(v.channelId)
	call = call.MaxResults(v.maxResults)
	call = call.PublishedAfter(v.publishAfter)
	call = call.PublishedBefore(v.publishBefore)
	if v.nextPageToken != "" {
		call = call.PageToken(v.nextPageToken)
	}

	// do it
	res, err := call.Do()
	if err != nil {
		return nil, errors.WithMessage(err, "data: GetVideos error: ")
	}
	if len(res.Items) > 0 {
		v.channelTitle = res.Items[0].Snippet.Title
	}
	v.nextPageToken = res.NextPageToken
	return res, err
}
