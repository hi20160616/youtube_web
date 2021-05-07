package json

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"google.golang.org/api/youtube/v3"
)

func TestActivitesList(t *testing.T) {
	tcs := []struct {
		cid  string
		want string
	}{
		{"UC_gUM8rL-Lrg6O3adPW9K1g", ""},
	}

	for _, tc := range tcs {
		ap := &ActivitiesParams{ChannelId: tc.cid, NextPageToken: "CBAQAA"}
		vs, err := &youtube.ActivityListResponse{}, errors.New("")
		i := 0
		for {
			vs, err = ap.
				WithMaxResults(10).
				WithPublishAfter(-1 * 60 * 24).
				List()
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

func TestGetAllChannelActivities(t *testing.T) {
	vl, err := getAllChannelActivities()
	if err != nil {
		t.Error(err)
	}
	for _, vs := range vl {
		for _, v := range vs.Items {
			fmt.Println(v.Snippet.Title)
		}
	}
}
