package server

import (
	"groupie-tracker/api"
	"strconv"
)

// CheckId verifies if the given id is a valid artist ID.
func CheckId(id string) bool {
	// Convert the id from string to integer
	value, err := strconv.Atoi(id)
	if err != nil {
		// Return false if conversion fails
		return false
	}
	// Get the total number of artists from the api
	len, err := api.GetArtistsLength()
	if err != nil {
		return false
	}
	// Check if the converted id is within the valid range
	if !(value > 0 && value <= len) {
		return false
	}
	// Return true if all checks pass
	return true
}
