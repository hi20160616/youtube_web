package render

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"time"

	"github.com/hi20160616/youtube_web/config"
	"github.com/rickb777/date/period"
)

var (
	tmplPath = filepath.Join(config.Value.TmplPath, "default")
)

type Page struct {
	Title string
	Data  interface{}
}

var templates = template.New("")

func init() {
	templates.Funcs(template.FuncMap{
		"summary":       Summary,
		"smartTime":     SmartTime,
		"smartLongTime": SmartLongTime,
		"smartDuration": SmartDuration,
	})
	pattern := filepath.Join(tmplPath, "*.html")
	templates = template.Must(templates.ParseGlob(pattern))
}

func Derive(w http.ResponseWriter, tmpl string, p *Page) {
	if err := templates.ExecuteTemplate(w, tmpl+".html", p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("err template: %s.html\n\terror: %#v", tmpl, err)
	}
}

func Summary(des string) string {
	dRune := []rune(des)
	if len(dRune) <= 300 {
		return des
	}
	return string(dRune[:300])
}

func parseWithZone(t string) time.Time {
	tt, err := time.Parse(time.RFC3339, t)
	if err != nil {
		log.Printf("render: SmartTime: %v", err)
	}
	loc := time.FixedZone("UTC", 8*60*60)
	return tt.In(loc)

}

func SmartTime(t string) string {
	return parseWithZone(t).Format("15:04/01.02")
}

func SmartLongTime(t string) string {
	return parseWithZone(t).String()
}

// https://play.golang.org/p/nMApT7G8SRV
func SmartDuration(t string) string {
	p, err := period.Parse(t)
	if err != nil {
		log.Printf("pkg: render: error: %v", err)
		return t
	}
	return fmt.Sprintf("%02d:%02d:%02d", p.Hours(), p.Minutes(), p.Seconds())
}
