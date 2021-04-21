package handler

import (
	"errors"
	"log"
	"net/http"
	"regexp"

	"github.com/hi20160616/youtube_web/internal/data"
	"github.com/hi20160616/youtube_web/internal/pkg/render"
	"google.golang.org/api/youtube/v3"
)

type Handler struct {
	YoutubeService *youtube.Service
	p              *render.Page
}

func NewHandler(s *youtube.Service) (*Handler, error) {
	return &Handler{YoutubeService: s}, nil
}

var validPath = regexp.MustCompile("^/(cid|vid|index|channels)/(.*?)$")

func (h *Handler) makeHandler(fn func(http.ResponseWriter, *http.Request, *render.Page)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, &render.Page{})
	}
}

func (h *Handler) GetHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		h.homeHandler(w, req)
		// fmt.Fprintf(w, "Welcome to the home page!")
	})
	mux.HandleFunc("/channels/", h.makeHandler(h.channelsHandler))
	mux.HandleFunc("/cid/", h.makeHandler(h.cidHandler))
	return mux
}

func (h *Handler) homeHandler(w http.ResponseWriter, r *http.Request) {
	// 1. get channels' ids that the VideoCount changed.
	// 2. get videos just uploaded
	// 3. render these videos
}

func (h *Handler) channelsHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	p.Title = "Channels"
	cs, err := data.ReadChannels()
	if err != nil {
		log.Printf("handler: channelsHandler: %v", err)
	}
	p.Data = cs
	render.Derive(w, "channels", p)
}

func (h *Handler) cidHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	// 1. GetVideos by channelIds
	cid := r.URL.Path[len("/cid/"):]
	vr := data.NewVideoRepo().SetChannelId(cid)
	res, err := &youtube.ActivityListResponse{}, errors.New("")
	res, err = vr.GetVideos()
	if err != nil {
		log.Printf("handler: cidHandler: %v", err)
	}
	if res.NextPageToken == "" {
		p.Data = "no videos"
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
	// p.Funcs = template.FuncMap{
	//         "summary":   render.Summary,
	//         "smartTime": render.SmartTime,
	// }
	// 3. render
	render.Derive(w, "cid", p)
}

func (h *Handler) vidHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	p.Title = "Video Title ?"
	// 1. GetVideo from api
	// 2. render it
}
