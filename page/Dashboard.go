package page

import (
	"log"
	"net/http"
	"text/template"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("template/dashboard.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = view.ExecuteTemplate(w, "dashboard.html", nil)
}
