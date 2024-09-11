package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

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

// GetWeatherData fetches weather data for a given zip code
func GetWeatherData(zipCode, apiKey string) (*WeatherResponse, error) {
	baseURL := "https://api.weatherapi.com/v1/current.json"
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	q := u.Query()
	q.Set("key", apiKey)
	q.Set("q", zipCode)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get weather data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var weatherResponse WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &weatherResponse, nil
}

// GetTemperatureCelsius returns the temperature in Celsius
func (w *WeatherResponse) GetTemperatureCelsius() float64 {
	return w.Current.TempC
}

// GetTemperatureFahrenheit returns the temperature in Fahrenheit
func (w *WeatherResponse) GetTemperatureFahrenheit() float64 {
	return w.Current.TempF
}

// GetWindSpeed returns the wind speed in mph
func (w *WeatherResponse) GetWindSpeed() float64 {
	return w.Current.WindMph
}

// GetWindDirection returns the wind direction
func (w *WeatherResponse) GetWindDirection() string {
	return w.Current.WindDir
}

// GetHumidity returns the humidity
func (w *WeatherResponse) GetHumidity() int {
	return w.Current.Humidity
}

// GetCloudCover returns the cloud cover
func (w *WeatherResponse) GetCloudCover() int {
	return w.Current.Cloud
}

// GetWeatherCondition returns the weather condition
func (w *WeatherResponse) GetWeatherCondition() string {
	return w.Current.Condition.Text
}

// GetCity returns the city name
func (w *WeatherResponse) GetCity() string {
	return w.Location.Name
}

// GetRegion returns the region name
func (w *WeatherResponse) GetRegion() string {
	return w.Location.Region
}

// GetCountry returns the country name
func (w *WeatherResponse) GetCountry() string {
	return w.Location.Country
}
