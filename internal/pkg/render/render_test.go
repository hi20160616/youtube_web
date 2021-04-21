package render

import (
	"testing"

	db "github.com/hi20160616/youtube_web/internal/pkg/db/json"
	"github.com/hi20160616/youtube_web/internal/pkg/handler"
)

func TestDerive(t *testing.T) {
	// cid := "UC-5VbWqa7FfpDaK2lLwE3dg"
	h, err := handler.NewHandler(db.YoutubeService)
	if err != nil {
		t.Error(err)
	}
	h.GetHandler()

}
