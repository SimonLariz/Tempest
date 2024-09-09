package location

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type zipCodeResponse struct {
	PostCode            string `json:"post code"`
	Country             string `json:"country"`
	CountryAbbreviation string `json:"country abbreviation"`
	Places              []struct {
		PlaceName         string `json:"place name"`
		Longitude         string `json:"longitude"`
		State             string `json:"state"`
		StateAbbreviation string `json:"state abbreviation"`
		Latitude          string `json:"latitude"`
	} `json:"places"`
}

func GetLocation() zipCodeResponse {
	// While the zip code is invalid
	for {
		// Get the zip code from the user
		fmt.Println("Enter a zip code: ")
		var zipCode string
		fmt.Scanln(&zipCode)

		// Validate the zip code
		zipCodeIsValid, zipCodeJSON := validateZipCode(zipCode)
		if zipCodeIsValid {
			// Print the location information
			getLocationInfo(zipCodeJSON)
			return zipCodeJSON
		} else {
			fmt.Println("Invalid zip code. Please try again.")
		}
	}
}

func validateZipCode(zipCode string) (bool, zipCodeResponse) {
	// Validate the zip code with the zippopotam.us API
	zipCodeAPIEndpoint := "https://api.zippopotam.us/us/" + zipCode

	resp, err := http.Get(zipCodeAPIEndpoint)
	if err != nil || resp.StatusCode != 200 {
		return false, zipCodeResponse{}
	}

	defer resp.Body.Close()

	// If the zip code is valid, create a zipCodeResponse struct
	var zipCodeResponse zipCodeResponse
	// Decode the response into the zipCodeResponse struct
	json.NewDecoder(resp.Body).Decode(&zipCodeResponse)

	return true, zipCodeResponse
}

func getLocationInfo(zipCodeJSON zipCodeResponse) {
	// Print the location information
	fmt.Println("Location: ")
	fmt.Println(zipCodeJSON.Places[0].PlaceName, ", ", zipCodeJSON.Places[0].State)
}
