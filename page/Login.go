package page

import (
	"log"
	"net/http"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("template/login.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = view.ExecuteTemplate(w, "login.html", nil)
}
