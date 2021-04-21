package render

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"time"
)

const (
	tmplPath = "templates/default"
)

type Page struct {
	Title string
	Data  interface{}
}

var templates = template.New("")

func init() {
	templates.Funcs(template.FuncMap{
		"summary":   Summary,
		"smartTime": SmartTime,
	})
	pattern := filepath.Join(tmplPath, "*.html")
	templates = template.Must(templates.ParseGlob(pattern))
}

func Derive(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Summary(des string) string {
	dRune := []rune(des)
	if len(dRune) <= 300 {
		return des
	}
	return string(dRune[:300])
}

func SmartTime(t string) string {
	tt, err := time.Parse(time.RFC3339, t)
	if err != nil {
		log.Printf("render: SmartTime: %v", err)
	}
	return tt.Format("[01.02][1504H]")
}
