package page

import (
	"log"
	"net/http"
	"text/template"
)

func Register(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("template/register.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = view.ExecuteTemplate(w, "register.html", nil)
}
