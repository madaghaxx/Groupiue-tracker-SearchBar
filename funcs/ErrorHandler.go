package funcs

import (
	"html/template"
	"net/http"
)

type ErrorPageData struct {
	Number  int
	Message string
}

func ErrorHandler(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	var message string
	switch code {
	case 500:
		message = "Internal Server Error"
	case 404:
		message = "Not Found"
	case 400:
		message = "Bad request"
	case 405:
		message ="Method Not allowed"
	}
	tmpl, err := template.ParseFiles("Templates/error.html")
	if err != nil {
		http.Error(w, "Unable to load error page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, ErrorPageData{Number: code, Message: message})
}
