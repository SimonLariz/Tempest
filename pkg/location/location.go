package location

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
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

func checkConfig() {
	// Check if the config file exists and create it if it doesn't
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		viper.Set("zip_code", "")
		viper.WriteConfigAs("config.yaml")
	}
}

func GetLocation() zipCodeResponse {
	// Check the config file
	checkConfig()
	// Check config file for ZIP code
	zipCode := viper.GetString("zip_code")
	if zipCode != "" {
		// Validate the zip code
		zipCodeIsValid, zipCodeJSON := validateZipCode(zipCode)
		if zipCodeIsValid {
			return zipCodeJSON
		}
	} else {
		// Get the zip code from the user
		zipCodeJSON := getZipCodeFromUser()
		return zipCodeJSON
	}
	return zipCodeResponse{}
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

func GetLocationInfo() {
	// Print the location information
	user_location := GetLocation()
	fmt.Println("Location: ")
	fmt.Println(user_location.Places[0].PlaceName, ", ", user_location.Places[0].State)
}

func GetZipCode() string {
	// Get the zip code from the user
	zipCodeJSON := GetLocation()
	return zipCodeJSON.PostCode
}

func getZipCodeFromUser() zipCodeResponse {
	// While the zip code is invalid
	for {
		// Get the zip code from the user
		fmt.Printf("Enter a zip code: ")
		var zipCode string
		fmt.Scanln(&zipCode)

		// Validate the zip code
		zipCodeIsValid, zipCodeJSON := validateZipCode(zipCode)
		if zipCodeIsValid {
			return zipCodeJSON
		} else {
			fmt.Println("Invalid zip code. Please try again.")
		}
	}
}

func promptUserToSaveZipCode(zip_code string) {
	// Prompt the user to save the zip code
	fmt.Printf("Would you like to save this zip code? (y/n): ")
	var saveZipCode string
	fmt.Scanln(&saveZipCode)

	if saveZipCode == "y" {
		// Save the zip code to the config file
		viper.Set("zip_code", zip_code)
		viper.WriteConfigAs("config.yaml")
	}
}

func ClearLocation() {
	// Clear the location
	viper.Set("zip_code", "")
	viper.WriteConfigAs("config.yaml")
}
