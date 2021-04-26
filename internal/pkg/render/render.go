package render

import (
	"fmt"
	"github.com/rickb777/date/period"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"time"
)

const (
	tmplPath = "templates/default" // for run
	// tmplPath = "../../../templates/default" // for testindex x 1
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
	return tt.Format("15:04/01.02")
}

// https://play.golang.org/p/nMApT7G8SRV
func SmartDuration(t string) string {
	p, err := period.Parse(t)
	if err != nil {
		log.Printf("pkg: render: error: %v", err)
		return t
	}
	return fmt.Sprintf("%02d:%02d:%02d", p.Hours(), p.Minutes(), p.Seconds())
	// t = strings.ToLower(t[2:])
	// tt, _ := time.ParseDuration(t)
	// ttt := tt.Truncate(time.Second).String()
	// r := strings.NewReplacer("h", ":", "m", ":", "s", "")
	// return r.Replace(ttt)
}
