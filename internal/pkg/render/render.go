package render

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"time"
)

const (
	tmplPath = "templates/default" // for run
	// tmplPath = "../../../templates/default" // for test
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
		"smartDuration": SmartDuration,
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

// https://play.golang.org/p/BbNjulPhw3m
func SmartDuration(t string) string {
	return t[2:]
	// t = strings.ToLower(t[2:])
	// tt, _ := time.ParseDuration(t)
	// ttt := tt.Truncate(time.Second).String()
	// r := strings.NewReplacer("h", ":", "m", ":", "s", "")
	// return r.Replace(ttt)
}
