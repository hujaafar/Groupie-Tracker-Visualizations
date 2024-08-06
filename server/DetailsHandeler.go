package server

import (
	"groupie-tracker/api"
	"html/template"
	"net/http"
)

// DetailsHandler handles the HTTP requests for the details page.
func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is GET.
	if r.Method != "GET" {
		renderErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the 'id' parameter from the URL query string.
	id := r.URL.Query().Get("id")
	// Validate the id using the CheckId function.
	isValid := CheckId(id)

	if !isValid {
		renderErrorPage(w, "Not Found", http.StatusNotFound)
		return
	}

	// Fetch data from different APIs using the id.
	data, err := api.GetAllDetails(id)
	if err != nil {
		renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Parse the HTML template file.
	tmpl, err := template.ParseFiles("static/html/details.html")
	if err != nil {
		renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Execute the template with the fetched data and write the output to the response writer.
	err = tmpl.Execute(w, data)
	if err != nil {
		renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
