package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/server"
)

func main() {
	// Set up the route for the home page.
	http.HandleFunc("/", server.HomeHandler)
	// Set up the route for the details page.
	http.HandleFunc("/details", server.DetailsHandler)
	// Serve static files from the "static" directory.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// Print a message to the console indicating that the server is running.
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
