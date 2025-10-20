package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

const (
	templateDir = "./templates/"
	base        = templateDir + "base.html"
	newGame     = templateDir + "new-game.html"
	game        = templateDir + "game.html"
	about       = templateDir + "about.html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, base, "index.html", nil)
}
func NewGame(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, base, "new-game.html", nil)
}
func Game(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, base, "game.html", nil)
}
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}
func renderTemplate(w http.ResponseWriter, base, page string, data any) {
	tpl := template.Must(template.ParseFiles(base, templateDir+page))
	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
