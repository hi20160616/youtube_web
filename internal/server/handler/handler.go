package handler

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/hi20160616/youtube_web/internal/data"
	db "github.com/hi20160616/youtube_web/internal/pkg/db/json"
	"github.com/hi20160616/youtube_web/internal/server/render"
	"google.golang.org/api/youtube/v3"
)

var validPath = regexp.MustCompile("^/(cid|cidNext|vid|index|channels|search)/(.*?)$")

type Handler struct {
	Channels   []*db.Channel
	Activities []*youtube.VideoListResponse
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
	mux.Handle("/s/", http.StripPrefix("/s/", http.FileServer(http.Dir("templates/default"))))
	mux.HandleFunc("/channels/", makeHandler(channelsHandler))
	mux.HandleFunc("/cid/", makeHandler(cidHandler))
	mux.HandleFunc("/cidNext/", makeHandler(cidNextHandler))
	mux.HandleFunc("/vid/", makeHandler(vidHandler))
	mux.HandleFunc("/search/", makeHandler(searchHandler))
	return mux
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 1. get activities videos by channelIds
	v, err := data.ReadActivities()
	if err != nil {
		log.Printf("homeHandler ReadActivities: %#v", err)
	}

	// 2. render these videos
	render.Derive(w, "index", &render.Page{Title: "Home", Data: v})
}

func channelsHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	h, err := data.ReadChannels()
	if err != nil {
		log.Printf("channelsHandler ReadChannels: %#v", err)
	}
	render.Derive(w, "channels", &render.Page{Title: "Channels", Data: h})
}

func cidHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	// 1. GetVideos by channelIds
	cid := r.URL.Path[len("/cid/"):]
	vr := data.NewVideosRepo().WithChannelId(cid).WithMaxResults(16)
	res, err := vr.List()
	if err != nil {
		log.Printf("handler: cidHandler: %v", err)
	}
	res.NextPageToken = vr.Activities.NextPageToken
	p.Data = res
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

func getCidAndToken(r *http.Request) (string, string) {
	return r.URL.Query().Get("cid"), r.URL.Query().Get("p")
}

func cidNextHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	cid, pToken := getCidAndToken(r)
	vr := data.NewVideosRepo().WithChannelId(cid).WithPageToken(pToken).WithMaxResults(16)
	res, err := vr.List()
	if err != nil {
		log.Printf("handler: cidNextHandler: %v", err)
	}
	res.NextPageToken = vr.Activities.NextPageToken
	p.Data = res
	// 2. Get channels title
	if p.Title == "" {
		if len(res.Items) > 0 {
			p.Title = res.Items[0].Snippet.ChannelTitle
		} else {
			p.Title = "No Title"
		}
	}
	// 3. render
	render.Derive(w, "videos", p)
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

func searchHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	p.Title = "Search"
	// get keywords
	u := r.URL.Query().Get("s")
	kws := strings.Split(u, " ")
	// search
	res, err := db.Search(kws...)
	if err != nil {
		log.Printf("handler: searchHandler: %#v", err)
	}
	if len(res.Items) > 0 {
		p.Data = res
	}
	// render
	render.Derive(w, "search", p)
}
