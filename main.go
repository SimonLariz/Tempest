package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

type weatherResponse struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		WindchillC float64 `json:"windchill_c"`
		WindchillF float64 `json:"windchill_f"`
		HeatindexC float64 `json:"heatindex_c"`
		HeatindexF float64 `json:"heatindex_f"`
		DewpointC  float64 `json:"dewpoint_c"`
		DewpointF  float64 `json:"dewpoint_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}

func loadEnvAPIKey() (string, error) {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	apiKey := os.Getenv("WEATHER_API_KEY")
	return apiKey, nil
}

func getZipCode() string {
	// Get the zip from user input with error handling
	var zipCode string
	// Loop until the user enters a valid zip code
	for {
		fmt.Print("Enter a zip code: ")
		fmt.Scanln(&zipCode)
		valid, zipCodeResp := validateZipCode(zipCode)
		if valid {
			fmt.Println("Entered zip for ", zipCodeResp.Places[0].PlaceName, ", ", zipCodeResp.Places[0].State, " is this correct? (y/n)")
			var confirm string
			fmt.Scanln(&confirm)
			return zipCode
		}
		fmt.Println("Invalid zip code. Please try again.")
	}

}

func validateZipCode(zipCode string) (bool, zipCodeResponse) {
	// Validate the zip code
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

func main() {
	// Get the API key from .env file
	apiKey, err := loadEnvAPIKey()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	fmt.Println("API Key: ", apiKey)
	userZipCode := getZipCode()
	fmt.Println("User Zip Code: ", userZipCode)
}
