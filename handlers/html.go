package handlers

import (
	"html/template"
	"net/http"
)

func UploadPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/upload.html")

	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
