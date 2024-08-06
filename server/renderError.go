package server

import (
	"html/template"
	"net/http"
	"strconv"
)

// renderErrorPage renders a custom error page with a given error message and status code.
func renderErrorPage(w http.ResponseWriter, errMsg string, statusCode int) {
	// Set the HTTP status code in the response header
	w.WriteHeader(statusCode)
	// Create an ErrorData struct with the error message and status code
	data := ErrorData{
		ErrorMessage: errMsg,
		Code:         strconv.Itoa(statusCode),
	}

	// Parse the HTML template for the error page
	tmpl, err := template.ParseFiles("static/html/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template with the ErrorData struct and write the output to the response writer
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
