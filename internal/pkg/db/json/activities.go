package json

import (
	"encoding/json"
	"io/ioutil"
	"os"
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

var ErrAcitvityListNil error = errors.New("ActivityListResponse is nil")

// var activitiesPath = "activities.json" // for test
var activitiesPath = "./db/activities.json"

func TimeParsed(minutes int) string {
	return time.Now().Add(time.Duration(minutes) * time.Minute).Format(time.RFC3339)
}

func (ap *ActivitiesParams) List() (*youtube.ActivityListResponse, error) {
	// init parameters
	if ap.Part == "" {
		ap.Part = "snippet,contentDetails"
	}

	if ap.ChannelId == "" {
		return nil, errors.New("db: ActivitiesParams List: cid is nil")
	}
	if ap.MaxResults == 0 {
		ap.MaxResults = 10
	}
	if ap.MaxResults >= 50 {
		ap.MaxResults = 50
	}
	if ap.PublishBefore == "" {
		ap.PublishBefore = TimeParsed(0)
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
	if ap.PageToken != "" {
		call = call.PageToken(ap.PageToken)
	}

	// do it
	res, err := call.Do()
	if err != nil {
		return nil, errors.WithMessage(err, "db: ActivitiesParams: List")
	}
	if len(res.Items) == 0 {
		return nil, ErrAcitvityListNil
	}
	ap.ChannelTitle = res.Items[0].Snippet.Title
	ap.NextPageToken = res.NextPageToken
	return res, err
}

func (ap *ActivitiesParams) WithChannelId(cid string) *ActivitiesParams {
	ap.ChannelId = cid
	return ap
}

// minutes should be a number, such as (-1 * 60 * 24) means 1 day ago from now on
func (ap *ActivitiesParams) WithPublishAfter(minutes int) *ActivitiesParams {
	if minutes == 0 {
		return ap
	}
	ap.PublishAfter = TimeParsed(minutes)
	return ap
}

func (ap *ActivitiesParams) WithMaxResults(max int64) *ActivitiesParams {
	ap.MaxResults = max
	return ap
}

func (ap *ActivitiesParams) WithPageToken(next string) *ActivitiesParams {
	ap.PageToken = next
	return ap
}

// getAllChannelActivities return VideoListResponse after 1 day ago
func getAllChannelActivities() ([]*youtube.VideoListResponse, error) {
	cs, err := ReadChannels()
	rt := []*youtube.VideoListResponse{}
	if err != nil {
		return nil, err
	}
	for _, c := range cs {
		vids := []string{}
		ap := &ActivitiesParams{ChannelId: c.ChannelId}
		activities, err := ap.WithPublishAfter(-1 * 24 * 60).List()
		if err != nil {
			if errors.Is(err, ErrAcitvityListNil) {
				continue
			}
			return nil, err
		}
		for _, vs := range activities.Items {
			if vs.ContentDetails.Upload != nil {
				vids = append(vids, vs.ContentDetails.Upload.VideoId)
			}
		}
		if len(vids) == 0 {
			continue
		}
		vp := &VideosParams{Id: strings.Join(vids, ",")}
		vls, err := vp.List()
		if err != nil {
			return nil, errors.WithMessage(err, "db activities err cid: "+c.ChannelId)
		}
		rt = append(rt, vls)

	}
	return rt, nil
}

func storageActivities(as []*youtube.VideoListResponse) error {
	asJson, err := json.Marshal(as)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(activitiesPath, asJson, 0644)
}

func ReadActivities() ([]*youtube.VideoListResponse, error) {
	return readActivities()
}

func readActivities() ([]*youtube.VideoListResponse, error) {
	as := []*youtube.VideoListResponse{}
	b, err := os.ReadFile(activitiesPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return UpdateActivities()
		} else {
			return nil, err
		}
	}
	if err = json.Unmarshal(b, &as); err != nil {
		return nil, err
	}
	return as, nil
}

func UpdateActivities() ([]*youtube.VideoListResponse, error) {
	as, err := getAllChannelActivities()
	if err != nil {
		return nil, err
	}
	if err = storageActivities(as); err != nil {
		return nil, err
	}
	return as, nil
}
