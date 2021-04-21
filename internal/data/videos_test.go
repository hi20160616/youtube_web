package data

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"google.golang.org/api/youtube/v3"
)

func TestGetVideos(t *testing.T) {
	res, err := &youtube.ActivityListResponse{}, errors.New("")

	vr := NewVideoRepo().SetChannelId("UC_gUM8rL-Lrg6O3adPW9K1g").SetMaxResults(10)
	i := 0
	for {
		res, err = vr.GetVideos()
		if err != nil {
			t.Error(err)
		}
		for _, v := range res.Items {
			i++
			fmt.Printf("%[2]s\t%[1]s\n", v.Snippet.Title, v.Snippet.PublishedAt)
		}
		fmt.Println(i)
		fmt.Println(len(res.Items))
		fmt.Println(res.PageInfo.TotalResults)
		if res.NextPageToken == "" {
			return
		}
		time.Sleep(5 * time.Second)
	}
}
