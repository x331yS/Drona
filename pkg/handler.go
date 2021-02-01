package pkg

import (
	"html/template"
	"net/http"
)

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/css/favicon.ico")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		switch r.Method {
		case "GET":
			t, err := template.ParseFiles("templates/index.html")
			if err != nil {
				InternalServerError(w, r)
			}
			_ = t.Execute(w, nil)
		case "POST":
			t, err := template.ParseFiles("templates/index.html")
			if err != nil {
				InternalServerError(w, r)
			}
			_ = t.Execute(w, nil)
		}

	} else {
		StatusNotFound(w, r)
	}
}
