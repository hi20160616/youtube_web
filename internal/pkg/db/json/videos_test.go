package json

import (
	"fmt"
	"testing"
)

func TestVideosList(t *testing.T) {
	tc := struct {
		id   string
		want string
	}{"Ks-_Mh1QhMc,c0KYU2j0TM4,eIho2S0ZahI", ""}
	vp := &VideosParams{Id: tc.id}
	vs, err := vp.List()
	if err != nil {
		t.Error(err)
	}
	for _, v := range vs.Items {
		fmt.Printf("title: %s [%s]\n", v.Snippet.Title, v.ContentDetails.Duration)
	}
}
