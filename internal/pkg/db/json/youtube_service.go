package json

import (
	"log"
	"net/http"
	"os/exec"
	"strings"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var YoutubeService *youtube.Service

func init() {
	newYoutubeService := func() (*youtube.Service, error) {
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
	s, err := newYoutubeService()
	if err != nil {
		log.Println(err)
	}
	YoutubeService = s
}
