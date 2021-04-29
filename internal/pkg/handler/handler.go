package handler

import (
	"log"
	"net/http"
	"regexp"

	"github.com/hi20160616/youtube_web/internal/data"
	db "github.com/hi20160616/youtube_web/internal/pkg/db/json"
	"github.com/hi20160616/youtube_web/internal/pkg/render"
	"google.golang.org/api/youtube/v3"
)

var validPath = regexp.MustCompile("^/(cid|vid|index|channels)/(.*?)$")

type Handler struct {
	Channels   []*db.Channel
	Activities []*youtube.VideoListResponse
}

var H = &Handler{nil, nil}

func init() {
	h, err := data.ReadChannels()
	if err != nil {
		log.Println(err)
	}
	v, err := data.ReadActivities()
	if err != nil {
		log.Println(err)
	}
	H.Channels = h
	H.Activities = v
}

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
	mux.HandleFunc("/vid/", makeHandler(vidHandler))
	return mux
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 1. get channels' ids that the VideoCount changed.
	// 2. get activities videos by channelIds
	p := &render.Page{Title: "Home", Data: H.Activities}

	// 3. render these videos
	render.Derive(w, "index", p)
}

func channelsHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	p.Title = "Channels"
	p.Data = H.Channels
	render.Derive(w, "channels", p)
}

func cidHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	// 1. GetVideos by channelIds
	cid := r.URL.Path[len("/cid/"):]
	vr := data.NewVideosRepo().WithChannelId(cid).WithMaxResults(24)
	// res, err := &youtube.VideoListResponse{}, errors.New("")
	res, err := vr.List()
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
	vid := r.URL.Path[len("/vid/"):]
	vr := data.NewVideosRepo().WithVideoId(vid)
	res, err := vr.GetVideos()
	if err != nil {
		log.Printf("handler: vidHandler: %v", err)
	}
	if len(res.Items) > 0 {
		p.Title = res.Items[0].Snippet.Title
		p.Data = res.Items[0]
	}
	// 2. render it
	render.Derive(w, "vid", p)
}
