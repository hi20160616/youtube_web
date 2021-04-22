package json

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"google.golang.org/api/youtube/v3"
)

func TestGetVideos(t *testing.T) {
	tcs := []struct {
		p    *Videos
		want string
	}{
		{&Videos{ChannelId: "UC_gUM8rL-Lrg6O3adPW9K1g", MaxResults: 10}, ""},
	}

	for _, tc := range tcs {
		vs, err := &youtube.ActivityListResponse{}, errors.New("")
		i := 0
		for {
			vs, err = tc.p.GetVideos()
			if err != nil {
				t.Error(err)
			}
			for _, v := range vs.Items {
				i++
				fmt.Printf("%[2]s\t%[1]s\n", v.Snippet.Title, v.Snippet.PublishedAt)
			}
			fmt.Println(i)
			fmt.Println(len(vs.Items))
			fmt.Println(vs.PageInfo.TotalResults)
			if vs.NextPageToken == "" {
				return
			}
			time.Sleep(5 * time.Second)
		}
	}
}
