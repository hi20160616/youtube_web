package data

import (
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"testing"
	"time"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func newYoutubeService() (*youtube.Service, error) {
	o, err := exec.Command("enit", "get", "yt_api").Output()
	if err != nil {
		return nil, err
	}
	api_key := string(strings.TrimSpace(strings.Split(string(o), "=")[1]))
	client := &http.Client{
		Transport: &transport.APIKey{Key: api_key},
	}
	return youtube.New(client)
}

func TestGetVideos(t *testing.T) {
	tcs := []struct {
		p    *Videos
		want string
	}{
		{&Videos{channelId: "UC_gUM8rL-Lrg6O3adPW9K1g", maxResults: 10}, ""},
	}

	s, err := newYoutubeService()
	if err != nil {
		t.Error(err)
	}

	for _, tc := range tcs {
		vs, err := &youtube.ActivityListResponse{}, errors.New("")
		i := 0
		for {
			vs, err = tc.p.GetVideos(s)
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
