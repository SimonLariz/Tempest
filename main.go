package main

import (
	"fmt"
	"os"

	"github.com/SimonLariz/Tempest/pkg/location"
	"github.com/joho/godotenv"
)

func loadEnvAPIKey() (string, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("API key not found in .env file, please add it")
	}
	return apiKey, nil
}

func main() {
	// Get the API key from .env file
	apiKey, err := loadEnvAPIKey()
	if err != nil {
		fmt.Println("Ensure you have a .env file with the WEATHER_API_KEY variable")
		return
	}
	fmt.Println("API Key: ", apiKey)

	// Get the location info
	zipCodeResp := location.GetLocation()
	zipCodeNumber := zipCodeResp.PostCode
	fmt.Println("Zip Code: ", zipCodeNumber)
}
