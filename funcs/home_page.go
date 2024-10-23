package funcs

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w,http.StatusMethodNotAllowed)
		return
	}

	fetchArtists()

	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	errr := tmpl.Execute(w, Artists)
	if errr != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

