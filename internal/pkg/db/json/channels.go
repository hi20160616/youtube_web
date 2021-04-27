package json

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/api/youtube/v3"
)

var (
	cidsPath string = "../../internal/pkg/db/json/cids.txt" // for test
	// cidsPath     string = "./internal/pkg/db/json/cids.txt" // for run
	channelsPath string = "./internal/pkg/db/json/channels.json"
)

type Channel struct {
	// snippet(.title), contentDetails(.uploads), statistics(.videoCount)
	Part         string // should be snippet,contentDetails
	ChannelId    string
	ChannelTitle string
	VideoCount   uint64 // should use for judgement channel videos updated
	Uploads      string // the uploads playlist
}

func getCids() ([]string, error) {
	f, err := os.Open(cidsPath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	ls := []string{}
	for scanner.Scan() {
		ls = append(ls, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return ls, nil
}

func GetChannelFromApi(s *youtube.Service, cid string) (*Channel, error) {
	c := &Channel{ChannelId: cid, Part: "snippet,contentDetails,statistics"}
	call := s.Channels.List(strings.Split(c.Part, ","))
	call = call.Id(c.ChannelId)

	// do it
	res, err := call.Do()
	if err != nil {
		return nil, errors.WithMessage(err, "GetChannel: ")
	}
	if len(res.Items) != 1 {
		return nil, errors.New("GetChannel: fetched err on ChannelId: " + c.ChannelId)
	}
	c.ChannelTitle = res.Items[0].Snippet.Title
	c.VideoCount = res.Items[0].Statistics.VideoCount
	c.Uploads = res.Items[0].ContentDetails.RelatedPlaylists.Uploads
	return c, nil
}

func GetChannels(s *youtube.Service) ([]*Channel, error) {
	cids, err := getCids()
	if err != nil {
		return nil, errors.WithMessage(err, "GetChannels: ")
	}
	cs := []*Channel{}
	for _, cid := range cids {
		c, err := GetChannelFromApi(s, cid)
		if err != nil {
			return nil, errors.WithMessage(err, "GetChannels: ")
		}
		cs = append(cs, c)
	}
	return cs, nil
}

func ReadChannels() ([]*Channel, error) {
	return readChannels()
}

func readChannels() ([]*Channel, error) {
	cs := []*Channel{}
	b, err := os.ReadFile(channelsPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return UpdateChannels()
		} else {
			return nil, err
		}
	}
	if err = json.Unmarshal(b, &cs); err != nil {
		return nil, err
	}
	return cs, nil
}

func UpdateChannels() ([]*Channel, error) {
	cs, err := GetChannels(YoutubeService)
	if err != nil {
		return nil, err
	}
	if err = storageChannels(cs); err != nil {
		return nil, err
	}
	return cs, nil
}

func storageChannels(cs []*Channel) error {
	csJson, err := json.Marshal(cs)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(channelsPath, csJson, 0644); err != nil {
		return nil
	}
	return nil
}
