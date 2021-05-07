package jobs

import (
	"fmt"
	"testing"
)

func TestUpdateActivities(t *testing.T) {
	if vs, err := UpdateActivities(); err != nil {
		t.Errorf("%#v", err)
	} else {
		for _, v := range vs {
			for _, vi := range v.Items {
				fmt.Println(vi.Snippet.Title)
			}
		}
	}
}
