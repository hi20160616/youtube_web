package json

import (
	"strings"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/api/youtube/v3"
)

// This struct is resumable with youtube api v3
type ActivitiesParams struct {
	Part          string // Always be snippet, contentDetails
	ChannelId     string
	ChannelTitle  string
	MaxResults    int64
	PageToken     string
	NextPageToken string
	PublishAfter  string
	PublishBefore string
	Fields        string
}

func timeParsed(minutes int) string {
	return time.Now().Add(time.Duration(minutes) * time.Minute).Format(time.RFC3339)
}

func (ap *ActivitiesParams) List() (*youtube.ActivityListResponse, error) {
	// init parameters
	if ap.Part == "" {
		ap.Part = "snippet,contentDetails"
	}

	if ap.ChannelId == "" {
		return nil, errors.New("db: json: cid is nil")
	}
	if ap.MaxResults == 0 {
		ap.MaxResults = 10
	}
	if ap.MaxResults >= 50 {
		ap.MaxResults = 50
	}
	if ap.PublishBefore == "" {
		ap.PublishBefore = timeParsed(0)
	}

	// load
	call := YoutubeService.Activities.List(strings.Split(ap.Part, ","))
	call = call.ChannelId(ap.ChannelId)
	call = call.MaxResults(ap.MaxResults)
	call = call.PublishedBefore(ap.PublishBefore)
	// call = call.Fields("contentDetails.upload.videoId")
	if ap.PublishAfter != "" {
		call = call.PublishedAfter(ap.PublishAfter)
	}
	if ap.NextPageToken != "" {
		call = call.PageToken(ap.NextPageToken)
	}

	// do it
	res, err := call.Do()
	if err != nil {
		return nil, errors.WithMessage(err, "db: activities: List")
	}
	if len(res.Items) > 0 {
		ap.ChannelTitle = res.Items[0].Snippet.Title
	}
	ap.NextPageToken = res.NextPageToken
	return res, err
}

// minutes should be a number, such as (-1 * 60 * 24) means 1 day ago from now on
func (ap *ActivitiesParams) WithPublishAfter(minutes int) *ActivitiesParams {
	ap.PublishAfter = timeParsed(minutes)
	return ap
}

func (ap *ActivitiesParams) WithMaxResults(max int64) *ActivitiesParams {
	ap.MaxResults = max
	return ap
}

// fields can be `contentDetails.upload.videoId` to fetch out videoId only
// func (ap *ActivitiesParams) WithFields(fields string) *ActivitiesParams {
//         ap.Fields = fields
//         return ap
// }
