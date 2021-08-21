package page

import (
	"log"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = view.ExecuteTemplate(w, "index.html", nil)
}
