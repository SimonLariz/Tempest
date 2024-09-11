package location

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

// GetLocationData fetches location data for a given zip code
func GetLocationData(zipCode string) (*zipCodeResponse, error) {
	baseURL := "https://api.zippopotam.us/us/"
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}

	u.Path += zipCode

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get location data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var zipCodeResponse zipCodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&zipCodeResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &zipCodeResponse, nil
}

// GetPostCode returns the post code
func (z *zipCodeResponse) GetPostCode() string {
	return z.PostCode
}

// GetCountry returns the country
func (z *zipCodeResponse) GetCountry() string {
	return z.Country
}

// GetCity returns the city
func (z *zipCodeResponse) GetCity() string {
	return z.Places[0].PlaceName
}

// GetState returns the state
func (z *zipCodeResponse) GetState() string {
	return z.Places[0].State
}
