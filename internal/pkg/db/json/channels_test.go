package json

import (
	"fmt"
	"testing"
)

func TestReadChannels(t *testing.T) {
	got, err := ReadChannels()
	if err != nil {
		t.Error(err)
	}
	if len(got) != 134 {
		t.Errorf("Read err, got: %d channels", len(got))
	}
	for _, v := range got {
		if v.ChannelTitle == "WION" {
			fmt.Println(v)
		}
	}
}

func TestGetChannel(t *testing.T) {
	tcs := []struct {
		c    *Channel
		want string
	}{
		{&Channel{ChannelId: "UC_gUM8rL-Lrg6O3adPW9K1g"}, "WION"},
	}

	for _, tc := range tcs {
		c, err := GetChannelFromApi(YoutubeService, tc.c.ChannelId)
		if err != nil {
			t.Error(err)
		}
		if tc.want != c.ChannelTitle || tc.want != tc.c.ChannelTitle {
			t.Errorf("want: %v\n got: %v", tc.want, c.ChannelTitle)
		}
		fmt.Println(tc.c.VideoCount)
	}
}

func TestGetChannels(t *testing.T) {
	cs, err := GetChannels(YoutubeService)
	if err != nil {
		t.Error(err)
	}
	for _, c := range cs {
		fmt.Println(c.ChannelTitle)
	}
}

// func TestSaveChannels(t *testing.T) {
//         s, err := newYoutubeService()
//         if err != nil {
//                 t.Error(err)
//         }
//         cs, err := data.GetChannels(s, "cids.txt")
//         if err != nil {
//                 t.Error(err)
//         }
//
//         if err = SaveChannels(cs, "../../json"); err != nil {
//                 t.Error(err)
//         }
// }
