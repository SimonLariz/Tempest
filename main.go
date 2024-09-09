package main

import (
	"fmt"
	"os"

	"github.com/SimonLariz/Tempest/pkg/location"
	"github.com/joho/godotenv"
)

func loadEnvAPIKey() (string, error) {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	apiKey := os.Getenv("WEATHER_API_KEY")
	return apiKey, nil
}

func main() {
	// Get the API key from .env file
	apiKey, err := loadEnvAPIKey()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	fmt.Println("API Key: ", apiKey)

	// Get the location info
	location.GetLocationInfo()
	zipCode := location.GetZipCode()
	fmt.Println("Zip Code: ", zipCode)
	location.ClearLocation()
}
