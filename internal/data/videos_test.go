package data

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"google.golang.org/api/youtube/v3"
)

func TestGetVideos(t *testing.T) {
	res, err := &youtube.VideoListResponse{}, errors.New("")

	vr := NewVideosRepo().
		WithChannelId("UC_gUM8rL-Lrg6O3adPW9K1g").
		WithMaxResults(10)
	i := 0
	for {
		res, err = vr.GetVideos()
		if err != nil {
			t.Error(err)
		}
		for _, v := range res.Items {
			i++
			fmt.Printf("%[2]s\t%[3]s\t%[1]s\n", v.Snippet.Title, v.Snippet.PublishedAt, v.ContentDetails.Duration)
		}
		fmt.Println(i)
		fmt.Println(len(res.Items))
		fmt.Println(res.PageInfo.TotalResults)
		if vr.Activities.NextPageToken == "" {
			return
		}
		time.Sleep(5 * time.Second)
	}
}
