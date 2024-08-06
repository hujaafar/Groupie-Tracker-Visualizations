package api

import (
	"encoding/json"
	"net/http"
)

// FetchLocation fetches location data for the given artist id.
func FetchLocation(id string) (LocationData, error) {
	// Make a GET request to the API endpoint for location data
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + id)
	if err != nil {
		return LocationData{}, err
	}
	defer res.Body.Close() // Ensure the response body is closed after the function returns

	// Decode the JSON response into a LocationData struct
	var data LocationData
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return LocationData{}, err
	}

	return data, nil
}

// FetchDetails fetches artist details for the given artist id.
func FetchDates(id string) (DatesData, error) {
	// Make a GET request to the API endpoint for artist details
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + id)
	if err != nil {
		return DatesData{}, err
	}
	defer res.Body.Close() // Ensure the response body is closed after the function returns

	// Decode the JSON response into a DatesData struct
	var data DatesData
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return DatesData{}, err
	}

	return data, nil
}

// FetchRelations fetches relation data for the given artist id.
func FetchRelations(id string) (RelationData, error) {
	// Make a GET request to the API endpoint for relation data
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		return RelationData{}, err
	}
	defer res.Body.Close() // Ensure the response body is closed after the function returns

	// Decode the JSON response into a RelationData struct
	var data RelationData
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return RelationData{}, err
	}

	return data, nil
}

// FetchDetails fetches artist details for the given artist id.
func FetchDetails(id string) (Details, error) {
	// Make a GET request to the API endpoint for artist details
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		return Details{}, err
	}
	defer resp.Body.Close() // Ensure the response body is closed after the function returns

	// Decode the JSON response into a Details struct
	var details Details
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		return Details{}, err
	}
	return details, nil
}

// GetAllDetails fetches all details (details, dates, relations, location) for the given artist id.
func GetAllDetails(id string) (AllDetails, error) {

	// Fetch artist details
	details, err := FetchDetails(id)
	if err != nil {
		return AllDetails{}, err
	}
	// Fetch dates data
	dates, err := FetchDates(id)
	if err != nil {
		return AllDetails{}, err
	}
	// Fetch relation data
	relation, err := FetchRelations(id)
	if err != nil {
		return AllDetails{}, err
	}
	// Fetch location data
	location, err := FetchLocation(id)
	if err != nil {
		return AllDetails{}, err
	}
	// Combine all the fetched data into an AllDetails struct
	allDetails := AllDetails{
		Details:   details,
		Dates:     dates,
		Relations: relation,
		Location:  location,
	}
	return allDetails, nil
}
