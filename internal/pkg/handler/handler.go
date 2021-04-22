package handler

import (
	"log"
	"net/http"
	"regexp"

	"github.com/hi20160616/youtube_web/internal/data"
	"github.com/hi20160616/youtube_web/internal/pkg/render"
)

var validPath = regexp.MustCompile("^/(cid|vid|index|channels)/(.*?)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, *render.Page)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, &render.Page{})
	}
}

func GetHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		homeHandler(w, req)
		// fmt.Fprintf(w, "Welcome to the home page!")
	})
	mux.HandleFunc("/channels/", makeHandler(channelsHandler))
	mux.HandleFunc("/cid/", makeHandler(cidHandler))
	return mux
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 1. get channels' ids that the VideoCount changed.
	// 2. get videos just uploaded
	// 3. render these videos
}

func channelsHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	p.Title = "Channels"
	cs, err := data.ReadChannels()
	if err != nil {
		log.Printf("handler: channelsHandler: %v", err)
	}
	p.Data = cs
	render.Derive(w, "channels", p)
}

func cidHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	// 1. GetVideos by channelIds
	cid := r.URL.Path[len("/cid/"):]
	vr := data.NewVideosRepo().WithChannelId(cid)
	// res, err := &youtube.VideoListResponse{}, errors.New("")
	res, err := vr.GetVideos()
	if err != nil {
		log.Printf("handler: cidHandler: %v", err)
	}
	if vr.Activities.NextPageToken == "" {
		p.Data = nil
	} else {
		p.Data = res
	}
	// 2. Get channels title
	if p.Title == "" {
		if len(res.Items) > 0 {
			p.Title = res.Items[0].Snippet.ChannelTitle
		} else {
			p.Title = "No Title"
		}
	}
	// 3. render
	render.Derive(w, "cid", p)
}

func vidHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	p.Title = "Video Title ?"
	// 1. GetVideo from api
	// 2. render it
}
