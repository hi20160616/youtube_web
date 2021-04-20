package render

import (
	"net/http"
	"path/filepath"
	"text/template"
)

const (
	tmplPath = "templates/default"
)

type Page struct {
	Title string
	Data  interface{}
	Funcs template.FuncMap
}

var templates = template.Must(template.ParseGlob(filepath.Join(tmplPath, "*.html")))

func Derive(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
