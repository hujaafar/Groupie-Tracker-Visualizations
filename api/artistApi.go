package api

import (
	"encoding/json"
	"net/http"
)

// FetchArtists fetches a list of artists from the API.
func FetchArtists() ([]Artist, error) {
	// Make a GET request to the API endpoint for artists
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // Ensure the response body is closed after the function returns

	// Decode the JSON response into a slice of Artist structs
	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}

	return artists, nil
}

// GetArtistsLength fetches the number of artists from the API.
func GetArtistsLength() (int, error) {
	// Fetch the list of artists using FetchArtists function
	artists, err := FetchArtists()
	if err != nil {
		return 0, err
	}
	// Return the length of the artists slice
	return len(artists), nil
}
