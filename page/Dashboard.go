package page

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("template/dashboard.html")
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(200)
	w.Write(dat)
}
