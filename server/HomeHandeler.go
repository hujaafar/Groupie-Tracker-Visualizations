package server

import (
	"groupie-tracker/api"
	"net/http"
	"text/template"
)

// ErrorData struct to hold error message and code for rendering error pages
type ErrorData struct {
	ErrorMessage string
	Code         string
}

// HomeHandler handles the HTTP requests for the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		renderErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Ensure the URL path is exactly "/". If not, return a "Not Found" error.
	if r.URL.Path != "/" {
		renderErrorPage(w, "Not Found", http.StatusNotFound)
		return
	}
	// Fetch the list of artists from the API.
	artists, err := api.FetchArtists()
	if err != nil {
		renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Parse the HTML template file for the home page.
	tmpl, err := template.ParseFiles("static/html/index.html")
	if err != nil {
		renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Execute the template with the fetched artists data and write the output to the response writer.
	err = tmpl.Execute(w, artists)
	if err != nil {
		renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
